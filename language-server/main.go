package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"quran-lsp/lsp"
	"quran-lsp/rpc"
	"quran-lsp/server"
	"strings"
	"unicode"
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
		// log.Println(state.Dump())
	})

	for scanner.Scan() {
		messageBytes := scanner.Bytes()
		response, err := HandleMessage(messageBytes, state)
		if err != nil {
			log.Printf("An error occured while creating the reponse, %v", err)
			continue
		}

		if response != nil {
			log.Println("Sending:", string(response))
		}

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

	log.Println("Recieved:", string(message))
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
	case "textDocument/didChange":
		var didChangeNotification lsp.DidChangeNotification
		if err := json.Unmarshal(message, &didChangeNotification); err != nil {
			return nil, err
		}
		state.SetDocument(
			didChangeNotification.Params.TextDocument.URI,
			didChangeNotification.Params.ContentChanges[0].Text,
		)
	case "textDocument/hover":
		var hoverMessage lsp.HoverMessage
		if err := json.Unmarshal(message, &hoverMessage); err != nil {
			return nil, err
		}
		document, found := state.GetDocument(hoverMessage.Params.TextDocument.URI)
		if found {
			documentLines := strings.Split(document, "\n")
			currentLine := documentLines[hoverMessage.Params.Position.Line]
			firstSpace := strings.IndexFunc(currentLine, func(r rune) bool { return unicode.IsSpace(r) })
			if firstSpace == -1 {
				firstSpace = len(currentLine)
			}
			token := currentLine[hoverMessage.Params.Position.Character:firstSpace]
			return rpc.Encode(
				lsp.NewHoverResponse(
					token,
					hoverMessage.Id,
				),
			)
		}
		return rpc.Encode(lsp.NewHoverResponse("", hoverMessage.Id))
	case "textDocument/completion":
		var completionMessage lsp.CompletionMessage
		if err := json.Unmarshal(message, &completionMessage); err != nil {
			return nil, err
		}
		return rpc.Encode(lsp.NewCompletionsItemResponse(completionMessage.Id, completionMessage.Params.Position))
	}
	return nil, nil
}
