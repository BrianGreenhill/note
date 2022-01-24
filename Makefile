.PHONY: all clean tools lint fmt
TOOLS=_tools

all: clean build install
	@clear
	@echo "`note`"

build: clean
	@go build -o ./bin/note main.go

install: all
	@echo "==> installing binary to ${HOME}/.local/bin"
	@cp ./bin/note $(HOME)/.local/bin

clean:
	@echo "==> cleaning generated files"
	@rm -rf ./bin/*

tools:
	@echo "==> installing tools"
	@mkdir -p $(TOOLS)
	@[ ! -f $(TOOLS)/go.mod ] && cd $(TOOLS) && go mod init tools || true
	@cd $(TOOLS) && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.30.0
	@cd $(TOOLS) && go install golang.org/x/tools/cmd/goimports@latest

fmt: tools
	@echo "==> cleaning imports"
	@goimports -local "github.com/BrianGreenhill/note" -l -w $(shell go list -f {{.Dir}} ./...)
	@echo "==> running go fmt"
	@gofmt -l -w .
	@echo "==> running go vet"
	@go vet .

lint: fmt
	@echo "==> running golang-ci-lint"
	@golangci-lint run ./...
