.PHONY: fmt lint

fmt:
		goimports -w .
		go fmt ./...

lint:
		golangci-lint run