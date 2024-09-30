#!/usr/bin/env bash

set -euo pipefail

docker compose exec quiz-app go test -race -v -count=1 -timeout 50s -coverpkg=./... -coverprofile=./tmp/coverage ./... | \
sed 's/===\s\+RUN/=== \o033[33mRUN\o033[0m/g' | \
sed 's/===/\o033[36m&\o033[0m/g' | \
sed 's/---/\o033[35m&\o033[0m/g' | \
sed '/PASS:/s/\(PASS:\s[^ ]*\s(\S*)\)/\o033[32m&\o033[0m/' | \
sed '/FAIL:/s/\(FAIL:\s[^ ]*\s(\S*)\)/\o033[31m&\o033[0m/'

docker compose exec quiz-app go tool cover -html=./tmp/coverage -o ./tmp/coverage.html
