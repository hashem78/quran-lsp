package lsp

type Notification struct {
	Message
	Id int `json:"id,omitempty"`
}
