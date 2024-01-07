.PHONY: test

test:
	@echo "##### Running tests"
	go test -race -cover -coverprofile=coverage -covermode=atomic -v ./...
