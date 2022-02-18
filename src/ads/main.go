package main

import (
	"src/ads/internal/conf"
	"src/ads/internal/infrastructure/ad"
	"src/ads/internal/logger"
	"src/ads/internal/transport/handlers"
	"src/ads/internal/transport/middlewares"
	"src/ads/internal/utils"
	"log"
	"net/http"
	"os"

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
	var adStorage ad.Storage

	switch config.Storage {
	case "cache":
		adStorage = ad.NewCache()
	case "database":
		adStorage = ad.NewSQL()
	}
	initRouter(config, adStorage)

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

func initRouter(config conf.Configuration, adStorage ad.Storage) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the ads API"))
	})
	router.Use(middlewares.AuthorizeService(config.Credentials))
	adRouter := router.PathPrefix("/ad").Subrouter()
	adRouter.HandleFunc("/", handlers.SearchAdsHandler(utils.SearchAds(adStorage))).Methods("GET")
	adRouter.Use(middlewares.AuthorizeSession(config.AuthService))
	adRouter.HandleFunc("/", handlers.CreateAdHandler(utils.CreateAd(adStorage))).Methods("POST")
	adRouter.HandleFunc("/{ad_id}", handlers.GetAdHandler(utils.GetAd(adStorage))).Methods("GET")
	adRouter.HandleFunc("/{ad_id}", handlers.UpdateAdHandler(utils.UpdateAd(adStorage))).Methods("PUT")
	adRouter.HandleFunc("/{ad_id}", handlers.DeleteAdHandler(utils.DeleteAd(adStorage))).Methods("DELETE")
	http.Handle("/", router)
}
