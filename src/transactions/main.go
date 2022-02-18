package main

import (
	"log"
	"net/http"
	"os"
	"src/transactions/internal/conf"
	"src/transactions/internal/infrastructure/transaction"
	"src/transactions/internal/logger"
	"src/transactions/internal/transport/handlers"
	"src/transactions/internal/transport/middlewares"
	"src/transactions/internal/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		panic("Missing .env file")
	}
	config, success := initConfig()

	if !success {
		os.Exit(1)
	}
	var transactionStorage transaction.Storage

	switch config.Storage {
	case "cache":
		transactionStorage = transaction.NewCache()
	case "database":
		//	transactionStorage = transaction.NewSQL()
	}
	initRouter(config, transactionStorage)

	server := &http.Server{Addr: config.Host + ":" + config.Port}
	logger.Log("Server listening on: " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func initConfig() (conf.Configuration, bool) {
	var config conf.Configuration
	file, err := os.Open("conf/" + os.Getenv("ENV") + ".yaml")

	if err != nil {
		logger.Log(err.Error())
		return config, false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Log("Cannot open config file")
		}
	}(file)

	decoder := yaml.NewDecoder(file)

	if err = decoder.Decode(&config); err != nil {
		logger.Log(err.Error())
		return config, false
	}
	return config, true
}

func initRouter(config conf.Configuration, transactionStorage transaction.Storage) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to the transactions API"))
		if err != nil {
			return
		}
	})
	router.Use(middlewares.AuthorizeService(config.Credentials))

	transactionRouter := router.PathPrefix("/transaction").Subrouter()
	transactionRouter.Use(middlewares.AuthorizeSession(config.AuthService))

	transactionRouter.HandleFunc("", handlers.GetAllHandler(utils.GetAll(transactionStorage))).Methods("GET")
	transactionRouter.HandleFunc("", handlers.CreateHandler(utils.Create(transactionStorage), handlers.GetAd(config.AdsService, config.AuthService))).Methods("POST")
	transactionRouter.HandleFunc("/{id}", handlers.GetHandler(utils.Get(transactionStorage))).Methods("GET")
	transactionRouter.HandleFunc("/{id}", handlers.UpdateHandler(utils.Update(transactionStorage))).Methods("PUT")
	transactionRouter.HandleFunc("/{id}", handlers.DeleteHandler(utils.Delete(transactionStorage))).Methods("DELETE")

	transactionRouter.HandleFunc("/{id}/accept", handlers.ActionHandler(utils.Action(transactionStorage, "accepted"))).Methods("PUT")
	transactionRouter.HandleFunc("/{id}/refuse", handlers.ActionHandler(utils.Action(transactionStorage, "refused"))).Methods("PUT")
	transactionRouter.HandleFunc("/{id}/cancel", handlers.ActionHandler(utils.Action(transactionStorage, "canceled"))).Methods("PUT")
	http.Handle("/", router)
}
