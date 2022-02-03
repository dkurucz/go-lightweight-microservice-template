package controllers

import (
	"fmt"
	"github.com/dkurucz/go-lightweight-microservice-template/src/models"
	"github.com/dkurucz/go-lightweight-microservice-template/src/utils"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"net/http"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// look up account by uuid
	account := models.FindAccountById(id)

	res, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := models.GetAllAccounts()
	res, _ := json.Marshal(accounts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// parse request
	accountReqData := &models.Account{}
	utils.ParseBody(r, accountReqData)
	
	// first ensure the account doesn't already exist by the specified name
	account := models.FindAccountByAccountName(accountReqData.AccountName)
	nullUuid, _ := uuid.FromString(models.NULL_ACCT_ID)
	if account.ID != nullUuid {
		http.Error(w, "Account already exists", 400)
		return
	}

	fmt.Println("Password:", accountReqData.Password)
	password := accountReqData.Password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	accountReqData.Password = string(bytes)

	// create a new account, one-way hashing the password, and setting timestamp
	accountReqData.CreateAccount()

	res, _ := json.Marshal(accountReqData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// parse request
	accountReqData := &models.Account{}
	utils.ParseBody(r, accountReqData)

	// first ensure the account doesn't already exist by the specified name
	account := models.FindAccountByAccountName(accountReqData.AccountName)
	nullUuid, _ := uuid.FromString(models.NULL_ACCT_ID)
	if account.ID == nullUuid {
		http.Error(w, "Account does not exist", 400)
		return
	}

	// parse account to JSON
	res, _ := json.Marshal(account)

	// delete account and write deleted account
	models.DeleteAccountById(account.ID)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
