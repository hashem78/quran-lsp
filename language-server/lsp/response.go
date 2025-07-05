package lsp

type Response struct {
	RPC   string         `json:"jsonrpc,omitempty"`
	Id    int            `json:"id,omitempty"`
	Error *ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    *any   `json:"data,omitempty"`
}
