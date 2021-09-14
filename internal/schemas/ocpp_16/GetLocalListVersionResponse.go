// Code generated by schema-generate. DO NOT EDIT.

package ocpp16

import (
    "bytes"
    "encoding/json"
    "fmt"
    "errors"
)

// GetLocalListVersionResponse 
type GetLocalListVersionResponse struct {
  ListVersion int `json:"listVersion"`
}

func (strct *GetLocalListVersionResponse) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "ListVersion" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "listVersion" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"listVersion\": ")
	if tmp, err := json.Marshal(strct.ListVersion); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetLocalListVersionResponse) UnmarshalJSON(b []byte) error {
    listVersionReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "listVersion":
            if err := json.Unmarshal([]byte(v), &strct.ListVersion); err != nil {
                return err
             }
            listVersionReceived = true
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    // check if listVersion (a required property) was received
    if !listVersionReceived {
        return errors.New("\"listVersion\" is required but was not present")
    }
    return nil
}
