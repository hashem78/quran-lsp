package lsp

type TextDocumentIdentifier struct {
	URI string `json:"uri,omitempty"`
}

type TextDocumentItem struct {
	TextDocumentIdentifier
	LanguageId string `json:"languageId,omitempty"`
	Version    int    `json:"version,omitempty"`
	Text       string `json:"text,omitempty"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
