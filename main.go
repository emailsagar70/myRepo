package main

import "github.com/gorilla/mux"

import "net/http"

import "github.com/tes-svc/handlers"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/account", handlers.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/v1/account", handlers.ListAccounts).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)

}
