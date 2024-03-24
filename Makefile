#!/usr/bin/env bash

.PHONY: skaffold
skaffold: ./bin/skaffold entgen swag
	@PATH="${PWD}/bin:${PATH}" ./bin/skaffold dev

./bin:
	mkdir bin

SUPPORT_SKAFFOLD_VERSION = latest
./bin/skaffold: ./bin
ifeq ($(shell uname), Darwin)
	@echo Downloading Skaffold ${SUPPORT_SKAFFOLD_VERSION} for MacOs
	curl -L -o ./bin/skaffold https://storage.googleapis.com/skaffold/releases/${SUPPORT_SKAFFOLD_VERSION}/skaffold-darwin-$(shell uname -m)
else
	@echo Please implement skaffold case for $(shell uname)
	@exit 1
endif
	@chmod +x ./bin/skaffold
	@./bin/skaffold version

./bin/mockery: ./bin
	GOBIN=${PWD}/bin go install github.com/vektra/mockery/v2@v2.42.1

./bin/swag: ./bin
	GOBIN=${PWD}/bin go install github.com/swaggo/swag/cmd/swag@v1.8.12

entgen:
	@cd api && go generate ./ent

mockery: ./bin/mockery
	@cd api && ../bin/mockery --dir=pkg --all

run-tests:
	@cd api && go test -v ./pkg/...

swag: ./bin/swag
	@cd api && ../bin/swag init --requiredByDefault=true --dir . --output ./docs -g **/**/*.go