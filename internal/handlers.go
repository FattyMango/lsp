package internal

import (
	"encoding/json"
	"lsp/models"
)
func (l *LSP) handleInitialize(content []byte) (any,error) {

	var req models.InitializeRequest
	if err := json.Unmarshal(content, &req); err != nil {
		l.logger.Println("Error unmarshalling initialize request:", err)
		return 	nil,err
	}

	l.logger.Println("Received initialize request:", req)

	return models.NewInitializeResponse(req.ID),nil

}

func (l *LSP) didOpen(content []byte) (any,error) {

	var req models.DidOpenTextDocumentNotification
	if err := json.Unmarshal(content, &req); err != nil {
		l.logger.Println("Error unmarshalling didOpen request:", err)
		return 	nil,err
	}

	l.logger.Println("Received didOpen request:", req)
	l.logger.Println("opened document:", req.Params.TextDocument.URI)

	l.state.OpenDocument(req.Params.TextDocument.URI, req.Params.TextDocument.Text)
	return nil,nil

}


func (l *LSP) didChange(content []byte) (any,error) {

	var req models.DidChangeTextDocumentNotification
	if err := json.Unmarshal(content, &req); err != nil {
		l.logger.Println("Error unmarshalling didChange request:", err)
		return 	nil,err
	}

	l.logger.Println("Received didChange request:", req)
	l.logger.Println("changed document:", req.Params.TextDocument.URI)
	l.logger.Println("new content:", req.Params.ContentChanges[0].Text)
	err := l.state.UpdateDocument(req.Params.TextDocument.URI, req.Params.ContentChanges[0].Text)
	if err != nil {
		l.logger.Println("Error updating document:", err)
		return nil,err
	}

	return nil,nil
}

func (l *LSP) hover (content []byte) (any,error) {
	var req models.HoverRequest
	if err := json.Unmarshal(content, &req); err != nil {
		l.logger.Println("Error unmarshalling hover request:", err)
		return 	nil,err
	}
	l.logger.Println("Received hover request:", req)
	return models.NewHoverResponse(req.ID),nil
}