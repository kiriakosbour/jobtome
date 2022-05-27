package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"urlShortener/controller"
	"urlShortener/domain"
	"urlShortener/service"

	"github.com/gorilla/mux"
)

type AppStruct struct {
}

func Start() {
	router := mux.NewRouter()
	urlRepo := domain.NewUrlRepoInit()
	urlService := service.UrlInteractorInit(urlRepo)
	addController := controller.UrlAddHandlerInit(urlService)
	getController := controller.UrlGetHandlerInit(urlService)
	allController := controller.UrlAllHandlerInit(urlService)
	router.HandleFunc("/api/create-short-url", addController.UrlAddController).Methods(http.MethodPost)
	router.HandleFunc("/api/{urlShortCode}", getController.UserGetController).Methods(http.MethodGet)
	router.HandleFunc("/api/count-all-url", allController.UrlAllController).Methods(http.MethodGet)
	address := os.Getenv("SERVER_ADDR")
	//address := "localhost"
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("No port given. Defaulting to port %s\n", port)
	}
	log.Print(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
