.PHONY: build dependencies check-style fmt vet lint

pkgs := $(shell go list ./...)

dependencies:
ifeq (, $(shell which revive))
	@echo "== cannot find revive installing"
	go install github.com/mgechev/revive@latest
endif
ifeq (, $(shell which fieldalignment))
	@echo "== cannot find fieldalignment installing"
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment
endif

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ccli-tz main.go

generate-mocks:
	mockgen -source=pkg/config/config.go -destination=pkg/config/mocks/config.go

check-style: dependencies fmt vet lint

.PHONY: check
check: check-style test

.PHONY: test
test:
	go test ./...

fmt:
ifneq (, $(shell gofmt -l .))
	$(error "gofmt found formatting issues: $(shell gofmt -l .). You may want to run `go fmt ./...` from the module folder")
endif

vet:
	go vet $(pkgs)
	go vet -vettool=$(which fieldalignment) $(pkgs)

lint:
	revive -config revive_conf.toml $(pkgs)
