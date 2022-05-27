package controller

import (
	"encoding/json"
	"net/http"
	"urlShortener/domain"
	"urlShortener/service"
)

type UrlAddHandler struct {
	service service.UrlInteractorInt
}

type UrlAddHandlerInterface interface {
	UrlAddController(w http.ResponseWriter, r *http.Request)
}

func UrlAddHandlerInit(adder service.UrlInteractorInt) *UrlAddHandler {
	return &UrlAddHandler{service: adder}
}
func (u *UrlAddHandler) UrlAddController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var urls domain.Urls
	errorMsg := json.NewDecoder(r.Body).Decode(&urls)
	if errorMsg != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMsg)
	}
	result := u.service.CreateTheUrlShortingService(urls.OriginalUrl)
	if result == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("There is an error during creation of the short url")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	}
}
