UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GORELEASER_BIN ?= https://github.com/goreleaser/goreleaser/releases/download/v1.7.0/goreleaser_Linux_x86_64.tar.gz
	REVIVE_BIN ?= https://github.com/mgechev/revive/releases/download/v1.2.0/revive_1.2.0_Linux_x86_64.tar.gz
endif
ifeq ($(UNAME_S),Darwin)
	GORELEASER_BIN ?= https://github.com/goreleaser/goreleaser/releases/download/v1.7.0/goreleaser_Darwin_all.tar.gz
	REVIVE_BIN ?= https://github.com/mgechev/revive/releases/download/v1.2.0/revive_1.2.0_Darwin_x86_64.tar.gz
endif

BOLD = \033[1m
CLEAR = \033[0m
CYAN = \033[36m

help: ## Display this help
	@awk '\
		BEGIN {FS = ":.*##"; printf "Usage: make $(CYAN)<target>$(CLEAR)\n"} \
		/^[a-z0-9]+([\/]%)?([\/](%-)?[a-z\-0-9%]+)*:.*? ##/ { printf "  $(CYAN)%-15s$(CLEAR) %s\n", $$1, $$2 } \
		/^##@/ { printf "\n$(BOLD)%s$(CLEAR)\n", substr($$0, 5) }' \
		$(MAKEFILE_LIST)

##@: Development

dev/clean:
	@rm -rf bin dist

dev/docs: bin/godoc ## Start the godoc server on :6060
	@bin/godoc -http :6060

##@: Test

test: ## Run the test suite
	@go test -cover ./... -p 8

test/coverage: bin/goverage ## Run tests with coverage report
	@bin/goverage -race -coverprofile=coverage.txt -covermode=atomic ./...

test/lint: bin/revive ## Lint go files
	@bin/revive -config revive.toml ./...

test/release: bin/goreleaser ## Create a local release snapshot
	@echo "$(CYAN)Creating snapshot build...$(CLEAR)"
	@bin/goreleaser --snapshot --rm-dist

##@: Binaries (local installations in ./bin)

bin/godoc: ## Install godoc
	$(call go-get-tool,$@,golang.org/x/tools/cmd/godoc)

bin/goreleaser: ## Install goreleaser
	@echo "$(CYAN)Installing goreleaser...$(CLEAR)"
	@mkdir -p bin
	@curl -sL $(GORELEASER_BIN) | tar xzf - -C bin
	@chmod +x bin/goreleaser
	@rm -rf bin/LICENSE.md bin/README.md bin/completions bin/manpages

bin/goverage: ## Install goverage
	$(call go-get-tool,$@,github.com/haya14busa/goverage)

bin/revive: ## Install revive
	@echo "$(CYAN)Installing revive...$(CLEAR)"
	@mkdir -p bin
	@curl -sL $(REVIVE_BIN) | tar xzf - -C bin
	@chmod +x bin/revive
	@rm -f bin/LICENSE bin/README.md

define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
echo "$(CYAN)Installing $(2)...$(CLEAR)" ;\
TMP_DIR=$$(mktemp -d) && cd $$TMP_DIR ;\
go mod init tmp && go get $(2) ;\
GOBIN=$(PWD)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
