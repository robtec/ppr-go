COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

DIST=dist

BIN=ppr

all: clean build

clean:
	@echo ">> cleaning..."
	@rm -rf $(DIST)

build:
	@go build -o dist/linux/ppr -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/ppr
	@go build -o dist/windows/ppr.exe -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/ppr
	@go build -o dist/osx/ppr -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/ppr

.PHONY: all clean build install