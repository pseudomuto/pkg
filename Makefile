.PHONY: coverage docs lint setup test

export GO111MODULE=on

TEST_PACKAGES = ./gen ./http

setup:
	# should match what's in tools.go
	@go install github.com/haya14busa/goverage
	@go install github.com/mgechev/revive

docs:
	@godoc -http :6060

lint:
	@revive -config revive.toml ./...

test:
	@go test -cover -race $(TEST_PACKAGES)

coverage:
	@goverage -race -coverprofile=coverage.txt -covermode=atomic $(TEST_PACKAGES)
