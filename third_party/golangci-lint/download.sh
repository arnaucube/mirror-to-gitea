#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
BINARY_PATH="$SCRIPT_PATH/golangci-lint"

DOWNLOAD_URL=''

if uname -a | grep 'Darwin' &> /dev/null; then
  DOWNLOAD_URL='https://github.com/golangci/golangci-lint/releases/download/v1.40.1/golangci-lint-1.40.1-darwin-amd64.tar.gz'
else
  DOWNLOAD_URL='https://github.com/golangci/golangci-lint/releases/download/v1.40.1/golangci-lint-1.40.1-linux-amd64.tar.gz'
fi

cd "$SCRIPT_PATH"
curl -fsSL "$DOWNLOAD_URL" | tar -xz

BINARY="$(find golangci-lint-*/golangci-lint)"
mv "$BINARY" "$BINARY_PATH"
rm -rf "$(find golangci-lint-*/ -type d)"
