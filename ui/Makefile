NPM=npm
MODE_INSTALL=$(NPM) install
BUILD=$(NPM) run build
BUILD_OPTIONS=NODE_OPTIONS=--openssl-legacy-provider
# NODE_OPTIONS=$(export NODE_OPTIONS=--openssl-legacy-provider;)
# # OS condition reference link: https://gist.github.com/sighingnow/deee806603ec9274fd47
# UNAME_S=$(shell uname -s)
# all:
# 	@echo $(OSFLAG)

# https://gist.github.com/sighingnow/deee806603ec9274fd47
# https://stackoverflow.com/questions/714100/os-detecting-makefile

ifeq ($(OS),Windows_NT)
	OS_ARCH += WIN32
	ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
		OS_ARCH += AMD64
	else
		ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
			OS_ARCH += AMD64
		endif
		ifeq ($(PROCESSOR_ARCHITECTURE),x86)
			OS_ARCH += IA32
		endif
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OS_ARCH += LINUX
	endif
	ifeq ($(UNAME_S),Darwin)
		OS_ARCH += OSX
	endif
	UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		OS_ARCH += AMD64
	endif
	ifneq ($(filter %86,$(UNAME_P)),)
		OS_ARCH += IA32
	endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		OS_ARCH += ARM
	endif
	ifneq ($(filter aarch%,$(UNAME_P)),)
		OS_ARCH += ARM
	endif
endif

ifeq ($(OS_ARCH),LINUX AMD)
	BUILD_KHOJI_UI=$(BUILD_OPTIONS) $(BUILD)
endif
ifeq ($(OS_ARCH),LINUX AMD64)
	BUILD_KHOJI_UI=$(BUILD_OPTIONS) $(BUILD)
endif
ifeq ($(OS_ARCH),LINUX ARM)
	BUILD_KHOJI_UI=$(BUILD_OPTIONS) $(BUILD)
endif
ifeq ($(OS_ARCH),OSX AMD)
	BUILD_KHOJI_UI=$(BUILD)
endif
ifeq ($(OS_ARCH),OSX IA32)
	BUILD_KHOJI_UI=$(BUILD)
endif
ifeq ($(OS_ARCH),OSX ARM)
	BUILD_KHOJI_UI=$(BUILD)
endif

all:
	$(MODE_INSTALL)
	@echo $(OS_ARCH)
	$(BUILD_KHOJI_UI)

clean:
	rm -rf dist
	rm -rf dist
	rm -rf node_modules
	rm -f package-lock.json
	rm -f yarn.lock