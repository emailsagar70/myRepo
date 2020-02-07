package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/tes-svc/models"
)

//LoginHandler is used to login to the account
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// this acc is created to hold the decoded json data
	var acc models.Account

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &acc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if acc.Email == "" || acc.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please enter correct Username and Password"))
		return
	}

	for i := range accounts {
		// Check if the email and password entered by user is correct or not
		if acc.Email == accounts[i].Email && acc.Password == accounts[i].Password {
			loginResponse := struct {
				LoginToken string `json:"logintoken"`
			}{
				LoginToken: uuid.New().String(),
			}

			resp, err := json.Marshal(&loginResponse)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			// Else condition is not written because  if 'if' condition doesn't match it will keep on going on else condition
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
			return

		}
	}
}
