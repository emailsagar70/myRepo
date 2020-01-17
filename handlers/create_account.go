package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tes-svc/models"
)

var (
	accounts []models.Account
)

//CreateAccount is used to create account details
func CreateAccount(w http.ResponseWriter, req *http.Request) {
	// this variable holds the decoded json data
	var account models.Account

	body, err := ioutil.ReadAll(req.Body)
	// fmt.Println(body)
	// fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // this will returnt the status code 400 when there is an error
		// User will receive the data in byte form this will convert the normal error in a string by decoding
		w.Write([]byte(err.Error()))
	}

	// It will decode all the data from json to code langugage
	if err = json.Unmarshal(body, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest) // this will returnt the status code 400 when there is an error
		// User will receive the data in byte form this will convert the normal error in a string by decoding
		w.Write([]byte(err.Error()))
	}

	// appending account to slice - accounts
	accounts = append(accounts, account)
	// after adding we have to return the status as created
	w.WriteHeader(http.StatusCreated)
	return

}
