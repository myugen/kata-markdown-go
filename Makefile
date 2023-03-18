.PHONY: download
download:
	@echo Download go.mod dependencies
	@go mod download

.PHONY: install-tools
install-tools: download
	@echo Installing tools from tools.go
	$(eval TOOLS := $(shell go list -f '{{range .Imports}}{{.}} {{end}}' tools.go))
	@go install $(TOOLS)

.PHONY: build
build: install-tools
	@echo Building link2footnote
	@go build -o link2footnote cmd/linkf2footnote/main.go
