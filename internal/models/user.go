// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"
	"time"

	"github.com/Beep-Technologies/beepbeep3-ocpp/pkg/constants"
)

// User is an object representing the database table.
type User struct {
	ID           int32     `gorm:"column:id"`
	Username     string    `gorm:"column:username"`
	Email        string    `gorm:"column:email"`
	HashedAPIKey string    `gorm:"column:hashed_api_key"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	Password     string    `gorm:"column:password"`
	IsAdmin      bool      `gorm:"column:is_admin"`
}

func (User) TableName() string {
	_ = strconv.Quote("")
	return constants.SCHEMA + "user"
}
