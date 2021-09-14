// Code generated by schema-generate. DO NOT EDIT.

package ocpp16

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// SendLocalListResponse 
type SendLocalListResponse struct {
  Status string `json:"status"`
}

func (strct *SendLocalListResponse) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Status" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "status" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *SendLocalListResponse) UnmarshalJSON(b []byte) error {
    statusReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "status":
            if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
                return err
             }
            statusReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if status (a required property) was received
    if !statusReceived {
        return errors.New("\"status\" is required but was not present")
    }
    return nil
}
