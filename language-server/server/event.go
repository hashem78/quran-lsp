package server

type EventType string

type StateEvent interface {
	GetType() EventType
}

const (
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
