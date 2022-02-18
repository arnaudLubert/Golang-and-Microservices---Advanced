package main

import (
    "src/authentication/internal/transport/middlewares"
    "src/authentication/internal/infrastructure/session"
    "src/authentication/internal/transport/handlers"
    "src/authentication/internal/logger"
    "src/authentication/internal/utils"
    "src/authentication/internal/conf"
    "github.com/joho/godotenv"
    "github.com/gorilla/mux"
    "gopkg.in/yaml.v2"
    "net/http"
    "log"
    "os"
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
    var sessionStorage session.Storage

    switch config.Storage {
    case "cache": sessionStorage = session.NewCache()
    case "database": sessionStorage = session.NewSQL()
    }
    initRouter(config, sessionStorage)

    server := &http.Server{ Addr: config.Host + ":" + config.Port }
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

func initRouter(config conf.Configuration, sessionStorage session.Storage) {
    router := mux.NewRouter()

    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Van Go API\n\nhttps://upload.wikimedia.org/wikipedia/commons/thumb/e/ea/Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg/757px-Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg"))
    })
    router.Use(middlewares.Authorize(config.Credentials))
    router.HandleFunc("/session", handlers.GetSessionHandler(utils.GetSession(sessionStorage, config.UsersService))).Methods("GET")
    router.HandleFunc("/login", handlers.LoginHandler(utils.CreateSession(sessionStorage), config.UsersService)).Methods("POST")
    router.HandleFunc("/logout", handlers.InvalidateSessionHandler(utils.InvalidateSession(sessionStorage))).Methods("GET")
    http.Handle("/", router)
}