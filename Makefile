.PHONY: docs lint test

export GO111MODULE=on

docs:
	@godoc -http :6060

lint:
	@revive -config revive.toml ./...

test:
	@go test -cover ./gen ./http
