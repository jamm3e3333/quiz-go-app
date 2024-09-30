#!/usr/bin/env bash

set -euo pipefail

if [ $# -ne 2 ]; then
    echo "Usage: $0 <test_pattern>"
    exit 1
fi
suite="$1"
pattern="$2"

go test -race -v -count=1 -run "$suite/$pattern" ./... | \
sed 's/===\s\+RUN/=== \o033[33mRUN\o033[0m/g' | \
sed 's/===/\o033[36m&\o033[0m/g' | \
sed 's/---/\o033[35m&\o033[0m/g' | \
sed '/PASS:/s/\(PASS:\s[^ ]*\s(\S*)\)/\o033[32m&\o033[0m/' | \
sed '/FAIL:/s/\(FAIL:\s[^ ]*\s(\S*)\)/\o033[31m&\o033[0m/'
