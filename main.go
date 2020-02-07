package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tes-svc/config"
	"github.com/tes-svc/db"
	"github.com/tes-svc/handlers"
)

func main() {
	// Establish a database connection
	db.NewClient()

	r := mux.NewRouter()
	r.HandleFunc("/v1/account", handlers.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/v1/account/login", handlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/v1/account", handlers.ListAccounts).Methods(http.MethodGet)
	r.HandleFunc("/v1/account/{id}", handlers.UpdateAccount).Methods(http.MethodPatch)
	r.HandleFunc("/v1/account/{id}", handlers.DeleteAccount).Methods(http.MethodDelete)

	http.ListenAndServe(":"+config.Port, r)

}
