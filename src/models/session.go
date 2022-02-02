package models

import (
	"github.com/dkurucz/go-lightweight-microservice-template/src/config"
	"github.com/jinzhu/gorm"
)

var sessionDb *gorm.DB

type Session struct {
	BaseModel
	Created         string  `json:"created" gorm:""`
	AccountName     string  `json:"account_name"`
	Expiration      string  `json:"expiration"`
}

func init() {
	sessionDb = config.GetDB()
}

func (session *Session) CreateSession() *Session {
	sessionDb.NewRecord(session)
	sessionDb.Create(&session)
	return session
}

func FindSessionById(ID string) *Session {
	var session Session
	sessionDb.Where("ID=?", ID).Find(&session)
	return &session
}

