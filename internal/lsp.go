package internal

import (
	"bufio"
	"log"
	"lsp/pkg/rpc"
	"lsp/pkg/state"
	"os"
)

type MessageHandler func(content []byte) (any, error)

type LSP struct {
	logger          *log.Logger
	messageHandlers map[string]MessageHandler
	state          *state.State
}

func NewLSP(logger *log.Logger) *LSP {
	l := &LSP{
		logger:          logger,
		messageHandlers: make(map[string]MessageHandler),
		state:          state.NewState(),
	}

	l.messageHandlers["initialize"] = l.handleInitialize
	l.messageHandlers["textDocument/didOpen"] = l.didOpen
	l.messageHandlers["textDocument/didChange"] = l.didChange
	l.messageHandlers["textDocument/hover"] = l.hover

	return l
}
func (l *LSP) Start() {

	l.logger.Println("LSP started")

	go l.listen()
}
func (l *LSP) Stop() {
	l.logger.Println("LSP stopped")
}
func (l *LSP) listen() {
	defer l.Stop()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			l.logger.Println("Error decoding message:", err)
			continue
		}
		l.handleMessage(method, content)
	}
}

func (l *LSP) handleMessage(method string, content []byte) {
	l.logger.Println("Received message:", method, string(content))

	handler, ok := l.messageHandlers[method]
	if !ok {
		l.logger.Println("No handler for method:", method)
		return
	}
	response, err := handler(content)
	if err != nil {
		l.logger.Println("Error handling message:", err)
		return
	}
	if response == nil {
		return
	}
	if err := l.reply(response); err != nil {
		l.logger.Println("Error replying to message:", err)
	}

}

func (l *LSP) reply(message any) error {

	writer := os.Stdout

	msg := rpc.EncodeMessage(message)

	l.logger.Println("Sending message:", string(msg))

	_, err := writer.Write([]byte(msg))

	return err

}
