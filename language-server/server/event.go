package server

type EventType string

type StateEvent interface {
	GetType() EventType
}

const (
	SERVER_START EventType = "server-start"
	ANY          EventType = "any"
	DOCUMENT_SET EventType = "document-set"
)

type DocumentSetEvent struct {
	Type EventType `json:"type,omitempty"`
	Uri  string    `json:"uri,omitempty"`
	Text string    `json:"text,omitempty"`
}

func (e DocumentSetEvent) GetType() EventType {
	return DOCUMENT_SET
}

type ServerStartEvent struct {
	Type EventType `json:"type,omitempty"`
}

func (e ServerStartEvent) GetType() EventType {
	return SERVER_START
}
