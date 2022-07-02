#!/bin/sh

set -eux

EXT=''
if [ $GOOS == 'windows' ]; then
EXT='.exe'
fi

make build

OUTPUT="${PROJECT_NAME}${EXT}"

echo ${OUTPUT}