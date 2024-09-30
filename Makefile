MIN_MAKE_VERSION := 3.81

# Min version
ifneq ($(MIN_MAKE_VERSION),$(firstword $(sort $(MAKE_VERSION) $(MIN_MAKE_VERSION))))
	$(error GNU Make $(MIN_MAKE_VERSION) or higher required)
endif

GO_LDFLAGS ?= -w -extldflags "-static" -X main.GitRevision=$(GIT_REVISION) -X main.Version=$(GIT_TAG_VERSION)
GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_TAG_VERSION := $(shell git tag -l --points-at HEAD | grep -v latest)

ifeq ($(CI),true)
	GO_TEST_EXTRAS ?= "-coverprofile=c.out"
endif

##@ Help
.PHONY: help
help: ## Show all available commands (you are looking at it)
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development
.PHONY: generate-proto build run up down exec-cli

generate-proto: ## Create grpc go files based on protobuffs
	script/generate-proto.sh

build: ## Build in docker
	./script/build.sh

run: up ## alias

up: ## Start up application container
	docker compose up --build

down: ## Stop and remove the application containers
	docker compose down --volumes

exec-cli: ## Go inside the application container
	docker compose exec quiz-app bash

##@ Test
.PHONY: test mtest mtest-pattern test-pattern test-suite-pattern
mtest: ## Run tests via docker-compose, supports arm arch
	./script/arm/test.sh

mtest-pattern: ## Run test with pattern, supports arm arch
	./script/arm/test-pattern.sh $(pattern)

test: ## Run tests via docker-compose
	./script/test.sh

test-pattern: ## Run test with pattern
	./script/test-pattern.sh $(pattern)

test-suite-pattern: ## Run test with pattern
	./script/test-suite-pattern.sh $(suite) $(pattern)

