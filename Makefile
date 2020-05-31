COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

DIST=dist

BIN=ppr

all: clean build

clean:
	@echo ">> cleaning..."
	@rm -rf $(DIST)

build:
	@go build -o ${DIST}/${BIN} -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID)" ./cmd/ppr

.PHONY: all clean build