COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

DIST=dist

BIN=ppr

all: clean build

clean:
	@echo ">> cleaning..."
	@rm -rf $(DIST)

install:
	@go install -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/ppr

.PHONY: all clean build install