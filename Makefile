.DIST_FOLDER := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build -o bin/countdown .
.PHONY:build

release: vet
	goreleaser release --snapshot --clean
.PHONY:release