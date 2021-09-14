// Code generated by schema-generate. DO NOT EDIT.

package ocpp16

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// ChargingProfile
type ChargingProfile struct {
	ChargingProfileId      int               `json:"chargingProfileId"`
	ChargingProfileKind    string            `json:"chargingProfileKind"`
	ChargingProfilePurpose string            `json:"chargingProfilePurpose"`
	ChargingSchedule       *ChargingSchedule `json:"chargingSchedule"`
	RecurrencyKind         string            `json:"recurrencyKind,omitempty"`
	StackLevel             int               `json:"stackLevel"`
	TransactionId          int               `json:"transactionId,omitempty"`
	ValidFrom              string            `json:"validFrom,omitempty"`
	ValidTo                string            `json:"validTo,omitempty"`
}

// RemoteStartTransactionRequest
type RemoteStartTransactionRequest struct {
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
	ConnectorId     int              `json:"connectorId,omitempty"`
	IdTag           string           `json:"idTag"`
}

func (strct *ChargingProfile) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// "ChargingProfileId" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfileId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfileId\": ")
	if tmp, err := json.Marshal(strct.ChargingProfileId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingProfileKind" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfileKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfileKind\": ")
	if tmp, err := json.Marshal(strct.ChargingProfileKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingProfilePurpose" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "chargingProfilePurpose" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfilePurpose\": ")
	if tmp, err := json.Marshal(strct.ChargingProfilePurpose); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "ChargingSchedule" field is required
	if strct.ChargingSchedule == nil {
		return nil, errors.New("chargingSchedule is a required field")
	}
	// Marshal the "chargingSchedule" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingSchedule\": ")
	if tmp, err := json.Marshal(strct.ChargingSchedule); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "recurrencyKind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"recurrencyKind\": ")
	if tmp, err := json.Marshal(strct.RecurrencyKind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "StackLevel" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "stackLevel" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stackLevel\": ")
	if tmp, err := json.Marshal(strct.StackLevel); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "transactionId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"transactionId\": ")
	if tmp, err := json.Marshal(strct.TransactionId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "validFrom" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"validFrom\": ")
	if tmp, err := json.Marshal(strct.ValidFrom); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "validTo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"validTo\": ")
	if tmp, err := json.Marshal(strct.ValidTo); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ChargingProfile) UnmarshalJSON(b []byte) error {
	chargingProfileIdReceived := false
	chargingProfileKindReceived := false
	chargingProfilePurposeReceived := false
	chargingScheduleReceived := false
	stackLevelReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingProfileId":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfileId); err != nil {
				return err
			}
			chargingProfileIdReceived = true
		case "chargingProfileKind":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfileKind); err != nil {
				return err
			}
			chargingProfileKindReceived = true
		case "chargingProfilePurpose":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfilePurpose); err != nil {
				return err
			}
			chargingProfilePurposeReceived = true
		case "chargingSchedule":
			if err := json.Unmarshal([]byte(v), &strct.ChargingSchedule); err != nil {
				return err
			}
			chargingScheduleReceived = true
		case "recurrencyKind":
			if err := json.Unmarshal([]byte(v), &strct.RecurrencyKind); err != nil {
				return err
			}
		case "stackLevel":
			if err := json.Unmarshal([]byte(v), &strct.StackLevel); err != nil {
				return err
			}
			stackLevelReceived = true
		case "transactionId":
			if err := json.Unmarshal([]byte(v), &strct.TransactionId); err != nil {
				return err
			}
		case "validFrom":
			if err := json.Unmarshal([]byte(v), &strct.ValidFrom); err != nil {
				return err
			}
		case "validTo":
			if err := json.Unmarshal([]byte(v), &strct.ValidTo); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if chargingProfileId (a required property) was received
	if !chargingProfileIdReceived {
		return errors.New("\"chargingProfileId\" is required but was not present")
	}
	// check if chargingProfileKind (a required property) was received
	if !chargingProfileKindReceived {
		return errors.New("\"chargingProfileKind\" is required but was not present")
	}
	// check if chargingProfilePurpose (a required property) was received
	if !chargingProfilePurposeReceived {
		return errors.New("\"chargingProfilePurpose\" is required but was not present")
	}
	// check if chargingSchedule (a required property) was received
	if !chargingScheduleReceived {
		return errors.New("\"chargingSchedule\" is required but was not present")
	}
	// check if stackLevel (a required property) was received
	if !stackLevelReceived {
		return errors.New("\"stackLevel\" is required but was not present")
	}
	return nil
}

func (strct *RemoteStartTransactionRequest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "chargingProfile" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"chargingProfile\": ")
	if tmp, err := json.Marshal(strct.ChargingProfile); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "connectorId" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"connectorId\": ")
	if tmp, err := json.Marshal(strct.ConnectorId); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "IdTag" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "idTag" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"idTag\": ")
	if tmp, err := json.Marshal(strct.IdTag); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RemoteStartTransactionRequest) UnmarshalJSON(b []byte) error {
	idTagReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "chargingProfile":
			if err := json.Unmarshal([]byte(v), &strct.ChargingProfile); err != nil {
				return err
			}
		case "connectorId":
			if err := json.Unmarshal([]byte(v), &strct.ConnectorId); err != nil {
				return err
			}
		case "idTag":
			if err := json.Unmarshal([]byte(v), &strct.IdTag); err != nil {
				return err
			}
			idTagReceived = true
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if idTag (a required property) was received
	if !idTagReceived {
		return errors.New("\"idTag\" is required but was not present")
	}
	return nil
}
