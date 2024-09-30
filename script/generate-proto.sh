#!/usr/bin/env bash
set -euo pipefail

echo "Generating go files based on protobuffs..."

protoc  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./grpc/protobuff/quiz.proto

echo "Go proto files have been generated."
