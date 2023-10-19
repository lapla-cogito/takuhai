BINDIR:=bin

ROOT_PACKAGE:=$(shell go list .)
COMMAND_PACKAGES:=$(shell go list ./...)

BINARIES:=$(COMMAND_PACKAGES:$(ROOT_PACKAGE)%=$(BINDIR)%)

build:
	go build -o $(BINDIR)/takuhai

format:
	go fmt ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
