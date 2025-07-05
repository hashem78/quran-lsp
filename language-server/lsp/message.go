package lsp

type Message struct {
	RPC    string `json:"jsonrpc,omitempty"`
	Method string `json:"method,omitempty"`
}
