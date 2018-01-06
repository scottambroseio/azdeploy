.PHONY: build install fmt test

GOFMT_FILES:=$$(find . -name '*.go' | grep -v vendor)

build: fmt
	go build -ldflags "-X github.com/scottrangerio/azuredeployment/cmd/version.version=0.1.0"  main.go

install: fmt
	go install -ldflags "-X github.com/scottrangerio/azuredeployment/cmd/version.version=0.1.0"  main.go

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	go test ./...
