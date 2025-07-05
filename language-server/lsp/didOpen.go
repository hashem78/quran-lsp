package lsp

type DidOpenNotification struct {
	Message
	Params DidOpenNotificationParams `json:"params"`
}

type DidOpenNotificationParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type TextDocumentItem struct {
	URI        string `json:"uri,omitempty"`
	LanguageId string `json:"language_id,omitempty"`
	Version    int    `json:"version,omitempty"`
	Text       string `json:"text,omitempty"`
}
