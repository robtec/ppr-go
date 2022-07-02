COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

DIST=dist

BIN=ppr

all: clean build

clean:
	@echo ">> cleaning..."
	@rm $(BIN)*

build:
	@go build -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID)" ./cmd/ppr

.PHONY: all clean build