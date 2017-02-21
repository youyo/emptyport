.DEFAULT_GOAL := help

## Setup
setup:
	go get github.com/Songmu/make2help/cmd/make2help

## Run tests
test: setup
	go test -cover

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup test help
