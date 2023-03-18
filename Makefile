download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	$(eval TOOLS := $(shell go list -f '{{range .Imports}}{{.}} {{end}}' tools.go))
	@go install $(TOOLS)

build: install-tools
	@echo Building link2footnote
	@go build -o link2footnote cmd/linkf2footnote/main.go
