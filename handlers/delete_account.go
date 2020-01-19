package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tes-svc/models"
)

// DeleteAccount will delete the details of an account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)

	var temp []models.Account

	// get all the account from the list
	for i := range accounts {

		if urlParams["id"] != accounts[i].ID {
			temp = append(temp, accounts[i])
		} else {
			fmt.Printf("Account %v is deleted successfully", urlParams["id"])
		}
	}

	accounts = temp
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("Account %v is deleted successfully", urlParams["id"])))

}
