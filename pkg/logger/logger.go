package logger

import (
	"log"
	"os"
)


func NewLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	return log.New(file, "[LSP]", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile)
}