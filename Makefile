FLUTTER_PLUGIN_PATH := lib/pigeon
FLUTTER_PLUGIN_SCHEME_FILE := $(FLUTTER_PLUGIN_PATH)/scheme.dart
FLUTTER_PLUGIN_GENERATE_FILE := $(FLUTTER_PLUGIN_PATH)/native.dart

ANDROID_PLUGIN_PATH := android/app/src/main/java/nagano/shunsuke/plugins
ANDROID_PLUGIN_PACAKGE :=  $(subst /,.,$(subst android/app/src/main/java/,,$(ANDROID_PLUGIN_PATH)))

ANDROID_PIGEON_FILE := $(ANDROID_PLUGIN_PATH)/Pigeon.java

IOS_PLUGIN_PATH := ios/Runner
IOS_PLUGIN_HEADER_FILE = $(IOS_PLUGIN_PATH)/Pigeon.h
IOS_PLUGIN_SOURCE_FILE = $(IOS_PLUGIN_PATH)/Pigeon.m
IOS_PLUGIN_PREFIX = Flutter

.PHONY: pigeon hoge

pigeon:
	mkdir -p $(ANDROID_PLUGIN_PATH)
	flutter pub run pigeon --input $(FLUTTER_PLUGIN_SCHEME_FILE) \
							--dart_out $(FLUTTER_PLUGIN_GENERATE_FILE) \
							--objc_header_out $(IOS_PLUGIN_HEADER_FILE) \
							--objc_source_out $(IOS_PLUGIN_SOURCE_FILE) \
							--objc_prefix $(IOS_PLUGIN_PREFIX) \
							--java_out $(ANDROID_PIGEON_FILE) \
							--java_package $(ANDROID_PLUGIN_PACAKGE)\
							--dart_null_safety
