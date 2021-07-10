.PHONY: run
run:
	@go run ./cmd/main.go

.PHONY: tests
tests:
	@go test -v -timeout 30s ./...

.PHONY: docs
docs:
	@godoc -http=localhost:9995