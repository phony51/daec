package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func getResultByID(w http.ResponseWriter, r *http.Request) {

}

func GetRouterResults() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/{id:[0-9]+}", getResultByID).Methods("GET")
	return router
}
