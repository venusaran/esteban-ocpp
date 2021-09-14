// Code generated by schema-generate. DO NOT EDIT.

package ocpp16

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// DataTransferRequest 
type DataTransferRequest struct {
  Data string `json:"data,omitempty"`
  MessageId string `json:"messageId,omitempty"`
  VendorId string `json:"vendorId"`
}

func (strct *DataTransferRequest) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "data" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"data\": ")
	if tmp, err := json.Marshal(strct.Data); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // Marshal the "messageId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"messageId\": ")
	if tmp, err := json.Marshal(strct.MessageId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "VendorId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "vendorId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"vendorId\": ")
	if tmp, err := json.Marshal(strct.VendorId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DataTransferRequest) UnmarshalJSON(b []byte) error {
    vendorIdReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "data":
            if err := json.Unmarshal([]byte(v), &strct.Data); err != nil {
                return err
             }
        case "messageId":
            if err := json.Unmarshal([]byte(v), &strct.MessageId); err != nil {
                return err
             }
        case "vendorId":
            if err := json.Unmarshal([]byte(v), &strct.VendorId); err != nil {
                return err
             }
            vendorIdReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if vendorId (a required property) was received
    if !vendorIdReceived {
        return errors.New("\"vendorId\" is required but was not present")
    }
    return nil
}
