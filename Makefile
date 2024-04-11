.PHONY: run
run:
	@go run cmd/main.go

.PHONY: build
build:
	@go build -o cmd/lsp cmd/main.go