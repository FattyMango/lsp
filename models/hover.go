package models

type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

type HoverParams struct {
	TextDocumentPositionParams
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type HoverResponse struct {
	Response
	Result HoverResult `json:"result"`
}

type HoverResult struct {
	Contents string `json:"contents"`
}

func NewHoverResponse(id int) *HoverResponse {
	return &HoverResponse{
		Response: Response{
			ID:  &id,
			RPC: "textDocument/hover",
		},
		Result: HoverResult{
			Contents: "Hello, From LSP!",
		},
	}
}
