package handlers

import (
	"net/http"

	"encoding/json"
)

// ListAccounts is used to get the list of an account
func ListAccounts(w http.ResponseWriter, r *http.Request) {
	// This will take the address of 'accounts' where all the 'Post'ed values are stored
	JSONResp, err := json.Marshal(&accounts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// if it gets the value, it'll set status = 200
	w.WriteHeader(http.StatusOK)
	// after response 200 it will send the value in bytes form to frontend.
	w.Write(JSONResp)
}
