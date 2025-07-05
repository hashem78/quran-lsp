package lsp

type RequestMessage struct {
	Message
	Id int `json:"id,omitempty"`
}
