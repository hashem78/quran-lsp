package lsp

type HoverMessage struct {
	Message
	Id     int                `json:"id,omitempty"`
	Params HoverMessageParams `json:"params"`
}

type HoverMessageParams struct {
	TextDocumentPositionParams
}

type HoverResponse struct {
	Response
	Result HoverResult `json:"result"`
}

type HoverResult struct {
	Contents string `json:"contents"`
	Range    *Range `json:"range"`
}

func NewHoverResponse(contents string, id int) HoverResponse {
	return HoverResponse{
		Response: Response{
			RPC: "2.0",
			Id:  id,
		},
		Result: HoverResult{
			Contents: contents,
			// Range: &Range{
			// 	Start: Position{
			// 		Line:      0,
			// 		Character: 0,
			// 	},
			// 	End: Position{
			// 		Line:      0,
			// 		Character: 2,
			// 	},
			// },
		},
	}
}
