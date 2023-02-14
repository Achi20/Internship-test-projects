package router

import (
	"go_test3/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	subrouter.HandleFunc("/{table}", middleware.GetAll).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/{table}/{id}", middleware.Get).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/new/{table}", middleware.Create2).Methods("POST", "OPTIONS")
	subrouter.HandleFunc("/{table}/{id}", middleware.Update).Methods("PUT", "OPTIONS")
	subrouter.HandleFunc("/delete/{table}/{id}", middleware.Delete).Methods("DELETE", "OPTIONS")

	return router
}
