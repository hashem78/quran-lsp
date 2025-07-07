package lsp

type DidChangeNotification struct {
	Message
	Params DidChangeNotificationParams `json:"params"`
}

type DidChangeNotificationParams struct {
	TextDocument   TextDocumentItem          `json:"textDocument"`
	ContentChanges []TextDocumentChangeEvent `json:"contentChanges"`
}

type TextDocumentChangeEvent struct {
	Text string `json:"text"`
}
