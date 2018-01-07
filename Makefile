.PHONY: build install fmt test vet cover

GO_FILES:=$$(find . -name '*.go' | grep -v vendor)

build:
	@go build -ldflags "-X github.com/scottrangerio/azuredeployment/cmd/version.version=0.1.0" azdeploy.go

install:
	@go install -ldflags "-X github.com/scottrangerio/azuredeployment/cmd/version.version=0.1.0" azdeploy.go

fmt:
	@gofmt -w $(GO_FILES)

vet:
	@go tool vet $(GO_FILES)

test:
	@go test ./cmd/...

cover:
	@go test ./cmd/... -cover
