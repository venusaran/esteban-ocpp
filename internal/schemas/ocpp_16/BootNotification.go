// Code generated by schema-generate. DO NOT EDIT.

package ocpp16

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// BootNotificationRequest 
type BootNotificationRequest struct {
  ChargeBoxSerialNumber string `json:"chargeBoxSerialNumber,omitempty"`
  ChargePointModel string `json:"chargePointModel"`
  ChargePointSerialNumber string `json:"chargePointSerialNumber,omitempty"`
  ChargePointVendor string `json:"chargePointVendor"`
  FirmwareVersion string `json:"firmwareVersion,omitempty"`
  Iccid string `json:"iccid,omitempty"`
  Imsi string `json:"imsi,omitempty"`
  MeterSerialNumber string `json:"meterSerialNumber,omitempty"`
  MeterType string `json:"meterType,omitempty"`
}

func (strct *BootNotificationRequest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "chargeBoxSerialNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"chargeBoxSerialNumber\": ")
	if tmp, err := json.Marshal(strct.ChargeBoxSerialNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ChargePointModel" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "chargePointModel" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"chargePointModel\": ")
	if tmp, err := json.Marshal(strct.ChargePointModel); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "chargePointSerialNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"chargePointSerialNumber\": ")
	if tmp, err := json.Marshal(strct.ChargePointSerialNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "ChargePointVendor" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "chargePointVendor" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"chargePointVendor\": ")
	if tmp, err := json.Marshal(strct.ChargePointVendor); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "firmwareVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"firmwareVersion\": ")
	if tmp, err := json.Marshal(strct.FirmwareVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "iccid" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"iccid\": ")
	if tmp, err := json.Marshal(strct.Iccid); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "imsi" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"imsi\": ")
	if tmp, err := json.Marshal(strct.Imsi); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "meterSerialNumber" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"meterSerialNumber\": ")
	if tmp, err := json.Marshal(strct.MeterSerialNumber); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "meterType" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"meterType\": ")
	if tmp, err := json.Marshal(strct.MeterType); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *BootNotificationRequest) UnmarshalJSON(b []byte) error {
    chargePointModelReceived := false
    chargePointVendorReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "chargeBoxSerialNumber":
            if err := json.Unmarshal([]byte(v), &strct.ChargeBoxSerialNumber); err != nil {
                return err
             }
        case "chargePointModel":
            if err := json.Unmarshal([]byte(v), &strct.ChargePointModel); err != nil {
                return err
             }
            chargePointModelReceived = true
        case "chargePointSerialNumber":
            if err := json.Unmarshal([]byte(v), &strct.ChargePointSerialNumber); err != nil {
                return err
             }
        case "chargePointVendor":
            if err := json.Unmarshal([]byte(v), &strct.ChargePointVendor); err != nil {
                return err
             }
            chargePointVendorReceived = true
        case "firmwareVersion":
            if err := json.Unmarshal([]byte(v), &strct.FirmwareVersion); err != nil {
                return err
             }
        case "iccid":
            if err := json.Unmarshal([]byte(v), &strct.Iccid); err != nil {
                return err
             }
        case "imsi":
            if err := json.Unmarshal([]byte(v), &strct.Imsi); err != nil {
                return err
             }
        case "meterSerialNumber":
            if err := json.Unmarshal([]byte(v), &strct.MeterSerialNumber); err != nil {
                return err
             }
        case "meterType":
            if err := json.Unmarshal([]byte(v), &strct.MeterType); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if chargePointModel (a required property) was received
    if !chargePointModelReceived {
        return errors.New("\"chargePointModel\" is required but was not present")
    }
    // check if chargePointVendor (a required property) was received
    if !chargePointVendorReceived {
        return errors.New("\"chargePointVendor\" is required but was not present")
    }
    return nil
}
