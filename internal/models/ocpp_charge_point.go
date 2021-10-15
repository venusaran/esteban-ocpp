// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/Beep-Technologies/beepbeep3-iam/pkg/constants"
)

// OcppChargePoint is an object representing the database table.
type OcppChargePoint struct {
	ID                      int32  `gorm:"column:id"`
	ApplicationID           int32  `gorm:"column:application_id"`
	ChargePointVendor       string `gorm:"column:charge_point_vendor"`
	ChargePointModel        string `gorm:"column:charge_point_model"`
	ChargePointSerialNumber string `gorm:"column:charge_point_serial_number"`
	ChargeBoxSerialNumber   string `gorm:"column:charge_box_serial_number"`
	Iccid                   string `gorm:"column:iccid"`
	Imsi                    string `gorm:"column:imsi"`
	MeterType               string `gorm:"column:meter_type"`
	MeterSerialNumber       string `gorm:"column:meter_serial_number"`
	FirmwareVersion         string `gorm:"column:firmware_version"`
	ConnectorCount          int32  `gorm:"column:connector_count"`
	ChargePointIdentifier   string `gorm:"column:charge_point_identifier"`
	OcppProtocol            string `gorm:"column:ocpp_protocol"`
}

func (OcppChargePoint) TableName() string {
	_ = strconv.Quote("")
	return constants.SCHEMA + "ocpp_charge_point"
}
