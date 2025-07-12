package lsp

type CompletionOptions struct {
	WorkDoneProgressOptions
}

type CompletionMessage struct {
	Message
	Id     int              `json:"id,omitempty"`
	Params CompletionParams `json:"params"`
}

type CompletionTriggerKind int

const (
	Invoked                         CompletionItemKind = 1
	TriggerCharacter                CompletionItemKind = 2
	TriggerForIncompleteCompletions CompletionItemKind = 3
)

type CompletionContext struct {
	TriggerKind      CompletionTriggerKind `json:"triggerKind,omitempty"`
	TriggerCharacter string                `json:"triggerCharacter,omitempty"`
}

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionItemsResponse struct {
	Response
	Result CompletionList `json:"result"`
}
type CompletionItemKind int

const (
	Text CompletionItemKind = 1
)

type InsertTextMode int

const (
	AsIs              InsertTextMode = 1
	AdjustIndentation InsertTextMode = 2
)

type CompletionList struct {
	Items        []CompletionItem `json:"items,omitempty"`
	IsIncomplete bool             `json:"isIncomplete,omitempty"`
}

type CompletionItem struct {
	Label          string             `json:"label,omitempty"`
	Kind           CompletionItemKind `json:"kind,omitempty"`
	Detail         string             `json:"detail,omitempty"`
	Documentation  string             `json:"documentation,omitempty"`
	InsertText     string             `json:"insertText,omitempty"`
	InsertTextMode InsertTextMode     `json:"insertTextMode,omitempty"`
	TextEdit       *TextEdit          `json:"textEdit,omitempty"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

func NewCompletionsItemResponse(id int, position Position) CompletionItemsResponse {
	return CompletionItemsResponse{
		Response: Response{
			RPC: "2.0",
			Id:  id,
		},
		Result: CompletionList{
			Items: []CompletionItem{
				{
					Label:  "Hashem",
					Kind:   1,
					Detail: "Inserting HelloWorld!",
				},
			},
			IsIncomplete: false,
		},
	}

}
