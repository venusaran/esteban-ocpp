// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/Beep-Technologies/beepbeep3-iam/pkg/constants"
)

// OcppApplicationCallback is an object representing the database table.
type OcppApplicationCallback struct {
	ID            int32  `gorm:"column:id"`
	ApplicationID int32  `gorm:"column:application_id"`
	CallbackEvent string `gorm:"column:callback_event"`
	CallbackURL   string `gorm:"column:callback_url"`
}

func (OcppApplicationCallback) TableName() string {
	_ = strconv.Quote("")
	return constants.SCHEMA + "ocpp_application_callback"
}
