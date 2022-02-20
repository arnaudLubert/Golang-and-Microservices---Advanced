package main

import (
	"log"
	"net/http"
	"os"
	"src/users/internal/conf"
	"src/users/internal/infrastructure/user"
	"src/users/internal/logger"
	"src/users/internal/transport/handlers"
	"src/users/internal/transport/middlewares"
	"src/users/internal/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	yaml "gopkg.in/yaml.v2"
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
	var userStorage user.Storage

	switch config.Storage {
	case "cache":
		userStorage = user.NewCache()
	case "database":
		userStorage = user.NewSQL()
	}
	initRouter(config, userStorage)

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
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err = decoder.Decode(&config); err != nil {
		logger.Log(err.Error())
		return config, false
	}
	return config, true
}

func initRouter(config conf.Configuration, userStorage user.Storage) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the users API"))
	})
	router.Use(middlewares.AuthorizeService(config.Credentials))
	router.HandleFunc("/auth/login", handlers.GetLoginHandler(utils.GetLogin(userStorage))).Methods("POST")
	router.HandleFunc("/auth/access/{user_id}", handlers.GetUserAccessHandler(utils.GetUserAccess(userStorage))).Methods("GET")
	router.HandleFunc("/new-account", handlers.CreateUserHandler(utils.CreateUser(userStorage))).Methods("POST")
	router.HandleFunc("/iban", handlers.GetUserIbanHandler(utils.GetUser(userStorage))).Methods("GET")

	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(middlewares.AuthorizeSession(config.AuthService))
	userRouter.HandleFunc("/{user_id}", handlers.GetUserHandler(utils.GetUser(userStorage))).Methods("GET")
	userRouter.HandleFunc("/", handlers.GetUsersHandler(utils.GetUsers(userStorage))).Methods("GET")
	userRouter.HandleFunc("/", handlers.UpdateUserHandler(utils.UpdateUser(userStorage))).Methods("PUT")
	userRouter.HandleFunc("/", handlers.DeleteUserHandler(utils.DeleteUser(userStorage))).Methods("DELETE")
	http.Handle("/", router)
}
