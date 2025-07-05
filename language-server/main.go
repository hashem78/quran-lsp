package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"quran-lsp/lsp"
	"quran-lsp/rpc"
	"quran-lsp/server"
)

func main() {

	file, err := os.OpenFile("/home/mythi/development/quran-lsp/lsp.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout

	state := server.NewServerState()
	state.AddListener(server.ANY, func(_ server.StateEvent) {
		log.Println(state.Dump())
	})

	for scanner.Scan() {
		messageBytes := scanner.Bytes()
		response, err := HandleMessage(messageBytes, state)
		if err != nil {
			log.Printf("An error occured while creating the reponse, %v", err)
			continue
		}

		// log.Println("Sending:", string(response))

		_, err = writer.Write(response)
		if err != nil {
			log.Printf("An error occured while writing the reponse, %v", err)
			continue
		}
	}
}

func HandleMessage(message []byte, state *server.State) ([]byte, error) {

	decodedMessage, err := rpc.Decode(message)

	if err != nil {
		log.Printf("An error occured while deocoding message, %v", err)
	}

	log.Println("Recieved:", decodedMessage.Method)
	switch decodedMessage.Method {
	case "initialize":
		var initializeMessage lsp.InitializeMessage
		if err := json.Unmarshal(message, &initializeMessage); err != nil {
			return nil, err
		}
		return rpc.Encode(lsp.NewInitializeResponse(initializeMessage.Id))
	case "initialized":
		var initializedNotification lsp.IntializedNotification
		if err := json.Unmarshal(message, &initializedNotification); err != nil {
			return nil, err
		}
	case "textDocument/didOpen":
		var didOpenNotification lsp.DidOpenNotification
		if err := json.Unmarshal(message, &didOpenNotification); err != nil {
			return nil, err
		}
		state.SetDocument(
			didOpenNotification.Params.TextDocument.URI,
			didOpenNotification.Params.TextDocument.Text,
		)
	}
	return nil, nil
}
