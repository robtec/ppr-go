#!/bin/sh

set -eux

COMMIT_ID=$(git rev-parse --short HEAD)
VERSION=$(cat VERSION)

go build -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID)" ./cmd/ppr

OUTPUT="${PROJECT_NAME}${EXT}"

echo ${OUTPUT}