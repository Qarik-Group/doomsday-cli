BUILD_TARGET ?= main.go
APP_NAME := doomsday
OUTPUT_NAME ?= $(APP_NAME)
SHELL := $(shell which bash)
VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X "main.appVersion=$(VERSION)"
BUILD := go build -v -ldflags='$(LDFLAGS)' -o $(OUTPUT_NAME) $(BUILD_TARGET)

.PHONY: build server darwin darwin-amd64 darwin-arm64 linux linux-amd64 all clean
.DEFAULT: build

build: server

server:
	@echo $(VERSION)
	GOOS=$(GOOS) GOARCH=$(GOARCH) VERSION=$(VERSION) $(BUILD)

darwin: darwin-amd64 darwin-arm64

darwin-amd64:
	GOOS=darwin GOARCH=amd64 OUTPUT_NAME=$(APP_NAME)-darwin-amd64 $(MAKE) server

darwin-arm64:
	GOOS=darwin GOARCH=arm64 OUTPUT_NAME=$(APP_NAME)-darwin-arm64 $(MAKE) server

linux: linux-amd64

linux-amd64:
	GOOS=linux GOARCH=amd64 OUTPUT_NAME=$(APP_NAME)-linux-amd64 $(MAKE) server

all: darwin linux

clean:
	rm -f $(APP_NAME) $(APP_NAME)-darwin-amd64 $(APP_NAME)-darwin-arm64 $(APP_NAME)-linux-amd64