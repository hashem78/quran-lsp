package rpc

import (
	"encoding/json"
	"quran-lsp/lsp"
)

func Decode(message []byte) (lsp.Message, error) {

	var decodedMessage lsp.Message
	err := json.Unmarshal(message, &decodedMessage)

	if err != nil {
		return lsp.Message{}, err
	}

	return decodedMessage, nil
}
