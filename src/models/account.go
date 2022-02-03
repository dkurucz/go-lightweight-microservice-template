package models

import (
	"github.com/dkurucz/go-lightweight-microservice-template/src/config"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

var accountDb *gorm.DB

type Account struct {
	BaseModel
	Created         time.Time   `json:"created" gorm:"autoCreateTime; column:created; default: NOW(); not null"`
	AccountName     string      `json:"account_name"`
	Password        string      `json:"password"`
}

func (acc *Account) CreateAccount() *Account {
	accountDb.NewRecord(acc)
	accountDb.Create(&acc)
	acc.Created = time.Now()
	return acc
}

func GetAllAccounts() []Account {
	var Accounts []Account
	accountDb.Find(&Accounts)
	return Accounts
}

func FindAccountById(id string) *Account {
	var account Account
	accountDb = config.GetDB()
	accountDb.Where("id=?", id).Find(&account)
	return &account
}

func FindAccountByAccountName(AccountName string) *Account {
	var account Account
	accountDb = config.GetDB()
	accountDb.Where("account_name=?", AccountName).Find(&account)
	return &account
}

func DeleteAccountById(id uuid.UUID) *Account {
	var account Account
	accountDb = config.GetDB()
	accountDb.Where("id=?", id.String()).Find(&account)
	accountDb.Unscoped().Delete(&account)
	return &account
}
