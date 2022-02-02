package models

import (
	"github.com/dkurucz/go-lightweight-microservice-template/src/config"
	"github.com/jinzhu/gorm"
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

func FindAccountById(ID string) *Account {
	var account Account
	accountDb = config.GetDB()
	accountDb.Where("ID=?", ID).Find(&account)
	return &account
}

func FindAccountByAccountName(AccountName string) *Account {
	var account Account
	accountDb = config.GetDB()
	accountDb.Where("account_name=?", AccountName).Find(&account)
	return &account
}

func DeleteAccountById(ID string) *Account {
	var account Account
	accountDb = config.GetDB()
	accountDb.Where("ID=?", ID).Find(&account)
	accountDb.Unscoped().Delete(&account)
	return &account
}
