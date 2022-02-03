package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

const NULL_ACCT_ID = "00000000-0000-0000-0000-000000000000"

type BaseModel struct {
	ID uuid.UUID  `gorm:"type:char(36);primary_key"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("id", uuid.NewV4())
}