.PHONY: lint
lint:
	golangci-lint run -v --fix

.PHONY: test
test:
	go test -covermode=count -coverprofile=count.out -v ./...
