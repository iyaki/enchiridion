.PHONY: help quality all format lint test coverage test-coverage test-race test-flaky \
	test-mutation mutation security arch build run release-check release-snapshot

BUILD_OUT ?= bin/enchiridion
ARGS ?=
FLAKY_COUNT ?= 20

help:
	@printf "%s\n" \
	"Common targets:" \
	"  make quality         Run full quality checks" \
	"  make format          Run gofmt on tracked Go files" \
	"  make lint            Run golangci-lint" \
	"  make test            Run tests" \
	"  make test-race       Run tests with race detector" \
	"  make test-flaky      Run tests repeatedly to detect flakes" \
	"  make coverage        Run coverage gate only" \
	"  make mutation        Run mutation testing (final stage only)" \
	"  make security        Run govulncheck and gosec" \
	"  make arch            Run go-arch-lint" \
	"  make build           Build CLI binary to $(BUILD_OUT)" \
	"  make run ARGS='...'  Run CLI from source" \
	"  make release-check    Validate goreleaser config" \
	"  make release-snapshot Build release artifacts locally (no publish)"

quality: test lint test-race test-flaky test-coverage test-mutation security arch

format:
	gofmt -w $$(git ls-files '*.go')

lint:
	golangci-lint run --timeout 5m

test:
	go test ./...

coverage: test-coverage

test-coverage:
	@if [ -z "$$(ls -d internal 2>/dev/null)" ]; then \
		echo "No internal packages yet — coverage gate deferred to phase 1."; \
		go test -count=1 ./... || exit 1; \
		exit 0; \
	fi; \
	coverprofile="$$(mktemp -t quality-cover.XXXXXX)"; \
	go test -count=1 -coverprofile="$$coverprofile" -covermode=atomic -coverpkg=./... ./...; \
	total="$$(go tool cover -func="$$coverprofile" | awk '/^total:/{gsub(/%/,"",$$3); print $$3}')"; \
	rm -f "$$coverprofile"; \
	if ! awk -v total="$$total" -v minimum="90" 'BEGIN {exit !(total >= minimum)}'; then \
		echo "Coverage $${total}% is below required 90%." >&2; \
		exit 1; \
	fi

test-race:
	go test -race ./...

test-flaky:
	go test -count=$(FLAKY_COUNT) -shuffle=on ./...

test-mutation:
	gremlins unleash $(ARGS)

mutation: test-mutation

security:
	govulncheck ./...
	gosec ./...

arch:
	go-arch-lint check

build:
	go build -o $(BUILD_OUT) ./cmd/enchiridion

run:
	go run ./cmd/enchiridion $(ARGS)

release-check:
	goreleaser check

release-snapshot:
	goreleaser release --snapshot --clean
