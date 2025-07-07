package lsp

type DidOpenNotification struct {
	Message
	Params DidOpenNotificationParams `json:"params"`
}

type DidOpenNotificationParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
