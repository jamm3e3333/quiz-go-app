#!/usr/bin/env bash

set -euo pipefail

GO_LDFLAGS=' -w -extldflags "-static"'
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

cd "$(dirname "$0")/.."
rm -rf target
mkdir -p target

echo "Building quiz-app..."

go build -ldflags "$GO_LDFLAGS" -o "target/server/quiz-app" -buildvcs=false "/go/src/github.com/jamm3e3333/quiz-app/cmd/start_server"
go build -ldflags "$GO_LDFLAGS" -o "target/cli/quiz-app" -buildvcs=false "/go/src/github.com/jamm3e3333/quiz-app/cmd/start_cli"

echo "Built: $(ls target/*)"
