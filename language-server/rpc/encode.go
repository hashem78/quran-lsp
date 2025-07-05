package rpc

import (
	"encoding/json"
	"fmt"
)

func Encode(contents any) ([]byte, error) {
	jsonBytes, err := json.Marshal(contents)
	if err != nil {
		return nil, err
	}
	response := fmt.Appendf(nil, "Content-Length: %d\r\n\r\n%s", len(jsonBytes), jsonBytes)
	return response, nil
}
