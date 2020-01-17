package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tes-svc/models"
)

var (
	// this creates a slice of account where it stores all the account details
	accounts []models.Account
)

//CreateAccount is used to create account details
func CreateAccount(w http.ResponseWriter, req *http.Request) {
	// this variable holds the decoded json data
	var account models.Account

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// this will return the status code 400 when there is an error
		w.WriteHeader(http.StatusBadRequest)
		// User will receive the data in byte form, this will convert the normal error into a string by decoding
		w.Write([]byte(err.Error()))
	}

	// It will decode all the data from json to code langugage
	if err = json.Unmarshal(body, &account); err != nil {
		// this will returnt the status code 400 when there is an error
		w.WriteHeader(http.StatusBadRequest)
		// User will receive the data in byte form, this will convert the normal error into a string by decoding
		w.Write([]byte(err.Error()))
	}

	// appending account to slice - accounts
	accounts = append(accounts, account)
	// after adding we have to return the status as created
	w.WriteHeader(http.StatusCreated)
	return

}
