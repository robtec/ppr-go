#!/bin/sh

set -eux

PROJECT_ROOT="/go/src/github.com/${GITHUB_REPOSITORY}"

mkdir -p $PROJECT_ROOT
rmdir $PROJECT_ROOT
ln -s $GITHUB_WORKSPACE $PROJECT_ROOT
cd $PROJECT_ROOT
go get -v ./...

COMMIT_ID=$(git rev-parse --short HEAD)
VERSION=$(cat VERSION)

EXT=''

if [ $GOOS == 'windows' ]; then
EXT='.exe'
fi

go build -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID)" ./cmd/ppr

OUTPUT="${PROJECT_NAME}${EXT}"

echo ${OUTPUT}