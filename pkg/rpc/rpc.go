package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {

	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(content []byte) (string, []byte, error) {
	header, content, ok := bytes.Cut(content, []byte("\r\n\r\n"))
	if !ok {
		return "", nil, fmt.Errorf("invalid message")
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}
	_ = content
	var msg BaseMessage
	err = json.Unmarshal(content[:contentLength], &msg)
	if err != nil {
		return "", nil, err
	}

	return msg.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {

	header, content, ok := bytes.Cut(data, []byte("\r\n\r\n"))
	if !ok {
		return 0, nil, fmt.Errorf("invalid message")
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
