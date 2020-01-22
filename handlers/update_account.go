package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tes-svc/models"
)

// UpdateAccount will update the details of the account
func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	// this variable holds the decoded json data
	var account models.Account

	// this will read the parameters requested in URL
	urlParams := mux.Vars(r)

	// This will check if there's an update or not
	var isUpdate bool

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	for i := range accounts {
		if urlParams["id"] == accounts[i].ID {
			if account.Firstname != "" {
				accounts[i].Firstname = account.Firstname
				isUpdate = true
			}
			if account.Middlename != "" {
				accounts[i].Middlename = account.Middlename
				isUpdate = true
			}
			if account.Lastname != "" {
				accounts[i].Lastname = account.Lastname
				isUpdate = true
			}
			if account.Email != "" {
				accounts[i].Email = account.Email
				isUpdate = true
			}
			if account.Phonenumber != "" {
				accounts[i].Phonenumber = account.Phonenumber
				isUpdate = true
			}
			if account.Password != "" {
				accounts[i].Password = account.Password
				isUpdate = true
			}
			if account.Gender != "" {
				accounts[i].Gender = account.Gender
				isUpdate = true
			}
			if account.Address != nil {
				if account.Address.Housenumber != "" {
					accounts[i].Address.Housenumber = account.Address.Housenumber
					isUpdate = true
				}
				if account.Address.Streetname != "" {
					accounts[i].Address.Streetname = account.Address.Streetname
					isUpdate = true
				}
				if account.Address.City != "" {
					accounts[i].Address.City = account.Address.City
					isUpdate = true
				}
				if account.Address.State != "" {
					accounts[i].Address.State = account.Address.State
					isUpdate = true
				}
				if account.Address.Zipcode != "" {
					accounts[i].Address.Zipcode = account.Address.Zipcode
					isUpdate = true
				}
			}
		}
	}
	// if there is no update then this will be displayed
	if isUpdate == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No changes detected"))
		return
	}

	// this is display after the update is being done
	w.WriteHeader(http.StatusNoContent)
	return

}
