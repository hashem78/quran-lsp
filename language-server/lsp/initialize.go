package lsp

type ClientInfo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type InitialzeMessageParams struct {
	ClientInfo ClientInfo `json:"clientInfo"`
}

type InitializeMessage struct {
	RequestMessage
	Params InitialzeMessageParams `json:"params"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync,omitempty"`
}

type ServerInfo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			Id:  id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
			},
			ServerInfo: ServerInfo{
				Name:    "quarn-lsp",
				Version: "0.0.0",
			},
		},
	}
}
