package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"urlShortener/domain"
	"urlShortener/service"
)

type UrlGetHandler struct {
	service service.UrlInteractorInt
}

type UrlGetHandlerInterface interface {
	UrlGetController(w http.ResponseWriter, r *http.Request)
}

func UrlGetHandlerInit(adder service.UrlInteractorInt) *UrlGetHandler {
	return &UrlGetHandler{service: adder}
}
func (u *UrlGetHandler) UserGetController(w http.ResponseWriter, r *http.Request) {
	var urls domain.Urls
	w.Header().Set("Content-Type", "application/json")
	shortUrl := mux.Vars(r)["urlShortCode"]
	urls.ShortUrl = shortUrl
	longUrl := u.service.RetrieveTheUrlShortingService(urls.ShortUrl)
	if longUrl == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error on retrieving the original url")
	} else {
		http.Redirect(w, r, "http://"+longUrl, 302)
	}
}
