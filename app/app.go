package app

import (
	lsp "lsp/internal"
	"lsp/pkg/logger"
	"os"
	"os/signal"
)

func Start() {
	logger := logger.NewLogger("/home/kassab/GO/lsp/lsp.log")
	LSP := lsp.NewLSP(logger)
	LSP.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
