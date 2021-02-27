PACKAGE_ROOT := github.com/glassmonkey/flutter_with_golang/golang/pkg
PACKAGE := compute
EXPORT_NAME := $(shell echo $(PACKAGE) | tr '[:lower:]' '[:upper:]')
.DEFAULT_GOAL := help
.PHONY: test  build-ios  build-android build help


test: ## run test
	go test -v ./...

build-ios: ## build go files for ios
	$(MAKE) build OS=ios EXPORT_NAME=$(EXPORT_NAME).framework

build-android: ## build go files for android
	#cp -r ~/Library/Android/sdk/ndk/22.0.7026061 ~/Library/Android/sdk/ndk-bundle
	ANDROID_NDK_HOME=~/Library/Android/sdk/ndk-bundle \
	$(MAKE) build OS=android EXPORT_NAME=$(EXPORT_NAME).aar

build:
	gomobile  bind  -target $(OS) -o ../$(OS)/$(EXPORT_NAME) $(PACKAGE_ROOT)/$(PACKAGE)

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
