GO := $(shell which go 2>/dev/null)
GO_LINT_WORKDIR ?= ""

.PHONY: help
help: ## Provides help information on available commands
	@printf "Usage: make <command>\n\n"
	@printf "Commands:\n"
	@awk -F ':(.*)## ' '/^[a-zA-Z0-9%\\\/_.-]+:(.*)##/ { \
	  printf "  \033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint
	@docker run --rm -v ${PWD}${GO_LINT_WORKDIR}:/app -w /app golangci/golangci-lint golangci-lint run --out-format tab | \
	awk -F '[[:space:]][[:space:]]+' '{ \
		error_file = $$1 ; \
		linter_name = $$2 ; \
		error_message = $$3 ; \
		split(error_file, error_file_info, ":") ; \
		error_file_path = sprintf(".%s/%s", ${GO_LINT_WORKDIR}, error_file_info[1]) ; \
		error_line_number = error_file_info[2] ; \
		error_col_number = error_file_info[3] ; \
		\
		dashed_line_length = 80 ; \
		dashed_line = sprintf("%*s", dashed_line_length, ""); gsub(/ /, "-", dashed_line) ; \
		\
		printf "\n\033[36m-- %s %0.*s %s\033[m", toupper(linter_name), dashed_line_length - length($$1) - length($$2), dashed_line, error_file ; \
		printf "\n\n\033[1mLine %s, Column %s", error_line_number, error_col_number ; \
		printf "\n\n\033[1m%s", error_message ; \
		\
		cmd_read_error_line = sprintf("sed -n %sp %s | sed -e \"s/\t/ /g\"", error_line_number, error_file_path) ; \
		cmd_read_error_line | getline error_line ; close(cmd_read_error_line) ; \
		printf "\n\n\033[33m%s|\033[m %s", error_line_number, error_line ; \
		\
		printf "\n\033[31m\033[1m%*s\033[m", error_col_number + length(error_line_number) + 2, "^" ; \
	} END { printf "\n\033[31m%s errors detected\n", NR	}'

.PHONY: fmt
fmt: $(GO) ## Format code according to Golang convention
	@go fmt ./...

.PHONY: test
test: $(GO) ## Run tests
	@-go test -race -covermode=atomic -coverprofile=coverage.out ./...

coverage.out:
	@$(MAKE) test

.PHONY: cover
cover: $(GO) coverage.out ## Given a coverage profile produced by 'go test'
	@go tool cover -func=coverage.out

.PHONY: cover/html
cover/html: $(GO) coverage.out ## Open a web browser that displays the coverage report
	@go tool cover -html=coverage.out

.PHONY: bench
bench: $(GO) ## Run benchmarks
	@go test -bench=. -benchmem ./...

.PHONY: gen/deepcopy
gen/deepcopy: $(GO) ## Generate a deep copy function
	@go run cmd/generator/main.go
