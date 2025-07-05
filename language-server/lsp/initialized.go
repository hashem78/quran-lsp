package lsp

type IntializedNotification struct {
	RequestMessage
	Params InitializedNotificationParams `json:"params"`
}

type InitializedNotificationParams struct {
}
