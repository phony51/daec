package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func addExpression(w http.ResponseWriter, r *http.Request) {

}

func getAllExpressions(w http.ResponseWriter, r *http.Request) {

}

func GetRouterExpressions() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", addExpression).Methods("POST")
	router.HandleFunc("/", getAllExpressions).Methods("GET")
	return router
}
