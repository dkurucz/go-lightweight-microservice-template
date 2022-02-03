package controllers

import (
	"github.com/dkurucz/go-lightweight-microservice-template/src/models"
	"github.com/dkurucz/go-lightweight-microservice-template/src/utils"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"

	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// process login
	var loginRequest map[string]interface{}
	loginRequest = utils.ParseBodyRaw(r)

	accountName := loginRequest["account_name"].(string)
	password := loginRequest["password"].(string)

	account := models.FindAccountByAccountName(accountName)
	nullUuid, _ := uuid.FromString(models.NULL_ACCT_ID)
	if account.ID == nullUuid {
		http.Error(w, "Invalid username or password", 403)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", 403)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login Successful"))
}