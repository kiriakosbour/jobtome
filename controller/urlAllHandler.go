package controller

import (
	"encoding/json"
	"net/http"
	"urlShortener/service"
)

type UrlAllHandler struct {
	service service.UrlInteractorInt
}

type UrlAllHandlerInterface interface {
	UrlAllController(w http.ResponseWriter, r *http.Request)
}

func UrlAllHandlerInit(adder service.UrlInteractorInt) *UrlAllHandler {
	return &UrlAllHandler{service: adder}
}
func (u *UrlAllHandler) UrlAllController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count := u.service.RetrieveAllTheUrlsSService()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("The total count of the shortened urls are: " + string(count))
	return
}
