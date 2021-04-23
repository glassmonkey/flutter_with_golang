PACKAGE_ROOT := github.com/kyloz/kyloz/golang/pkg
PACKAGE := example
EXPORT_NAME := $(shell echo $(PACKAGE) | tr '[:lower:]' '[:upper:]')
ANDROID_NDK_VERSION := 22.0.7026061

.DEFAULT_GOAL := help
.PHONY: test build-all build-ios build-android build help

test: ## run test
	go test -v ./...

build-all:
	$(MAKE) build-android
	$(MAKE) build-ios

build-ios: ## build go files for ios
	$(MAKE) build OS=ios EXPORT_NAME=$(EXPORT_NAME).framework

build-android: ## build go files for android
	ANDROID_NDK_HOME=~/Library/Android/sdk/ndk/$(ANDROID_NDK_VERSION) \
	$(MAKE) build OS=android EXPORT_NAME=$(EXPORT_NAME).aar

build:
	gomobile bind -target $(OS) -o ../$(OS)/$(EXPORT_NAME) $(PACKAGE_ROOT)/$(PACKAGE)

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
