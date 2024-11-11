.PHONY:lint
lint:
	golangci-lint run --fix

.PHONY: test
test:
	go test ./...

.PHONY: install
install:
	go install ./
