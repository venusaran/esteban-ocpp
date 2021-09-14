// Code generated by schema-generate. DO NOT EDIT.

package ocpp16

import (
    "encoding/json"
    "fmt"
    "bytes"
)

// GetDiagnosticsResponse 
type GetDiagnosticsResponse struct {
  FileName string `json:"fileName,omitempty"`
}

func (strct *GetDiagnosticsResponse) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // Marshal the "fileName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"fileName\": ")
	if tmp, err := json.Marshal(strct.FileName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GetDiagnosticsResponse) UnmarshalJSON(b []byte) error {
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "fileName":
            if err := json.Unmarshal([]byte(v), &strct.FileName); err != nil {
                return err
             }
        default:
            return fmt.Errorf("additional property not allowed: \"" + k + "\"")
        }
    }
    return nil
}
