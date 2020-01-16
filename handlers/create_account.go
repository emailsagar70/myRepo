package handlers

import (
	"github.com/tes-svc/models"
	"io/ioutil"
	"net/http"
)

var (
	accounts []models.Account
)

//CreateAccount is used to create account details
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// this variable holds the decoded jason data
	var account models.Account

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // this will returnt the status code 400 when there is an error
		// User will receive the data in byte form this will convert the normal error in a string by decoding
		w.Write([]byte(err.Error()))
	}
}
