# Go parameters
GOCMD=go
GOPATH=$(shell go env GOPATH)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
#GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CGO_CFLAGS=$(shell env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/")
CGO_LDFLAGS_DARWIN=$(shell env CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security")
CGO_LDFLAGS_WIN="-L$(GOPATH)/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv"
CGO_LDFLAGS_LINUX="-L$(GOPATH)/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm"
CGO_CC_WIN="x86_64-w64-mingw32-gcc"
MKDIR_P=mkdir -p
GITCMD=git
ROOT_DIR=$(shell pwd)
CREATEDB_DIR=$(ROOT_DIR)/createdb/
BINARY_NAME=khoji
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_OSX=$(BINARY_NAME)_osx
BINARY_WIN=$(BINARY_NAME).exe
SETUP_BINARY_NAME=createdb
SETUP_BINARY_UNIX=$(SETUP_BINARY_NAME)_unix
SETUP_BINARY_OSX=$(SETUP_BINARY_NAME)_osx
SETUP_BINARY_WIN=$(SETUP_BINARY_NAME).exe
DIST_DIR=dist
DIST_OSX=khoji_osx
DIST_OSX_PATH=$(DIST_DIR)/$(DIST_OSX)
DIST_OSX_ARM=khoji_osx_arm
DIST_OSX_ARM_PATH=$(DIST_DIR)/$(DIST_OSX_ARM)
DIST_UNIX=khoji_unix
DIST_UNIX_PATH=$(DIST_DIR)/$(DIST_UNIX)
DIST_WIN=khoji_win
DIST_WIN_PATH=$(DIST_DIR)/$(DIST_WIN)
DIST_FILES=LICENSE README.md
CP_AV=cp -av
CURL_DL=curl -LJ
RM_RFV=rm -rfv
UNZIP=unzip
TAR_GZ=tar -cvzf
CHECKOUT_BRANCH=main
DEPS_LINUX=GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_CFLAGS="-I$(GOPATH)/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(GOPATH)/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm"
DEPS_OSX=GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-I$(GOPATH)/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(GOPATH)/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security"
DEPS_OSX_ARM=GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS="-I$(GOPATH)/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(GOPATH)/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security"
DEPS_WIN=GO111MODULE=auto CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_CFLAGS="-I$(GOPATH)/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(GOPATH)/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc"

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
	BUILD_DEPS=deps-linux
	BUILD_KHOJI=$(DEPS_LINUX) $(GOBUILD) -o $(BINARY_NAME) -v
endif
ifeq ($(OS_ARCH),LINUX AMD64)
	BUILD_DEPS=deps-linux
	BUILD_KHOJI=$(DEPS_LINUX) $(GOBUILD) -o $(BINARY_NAME) -v
endif
ifeq ($(OS_ARCH),LINUX ARM)
	BUILD_DEPS=deps-linux
	BUILD_KHOJI=$(DEPS_LINUX) $(GOBUILD) -o $(BINARY_NAME) -v
endif
ifeq ($(OS_ARCH),OSX AMD)
	BUILD_DEPS=deps-osx
	BUILD_KHOJI=$(DEPS_OSX) $(GOBUILD) -o $(BINARY_NAME) -v
endif
ifeq ($(OS_ARCH),OSX IA32)
	BUILD_DEPS=deps-osx
	BUILD_KHOJI=$(DEPS_OSX) $(GOBUILD) -o $(BINARY_NAME) -v
endif
ifeq ($(OS_ARCH),OSX ARM)
	BUILD_DEPS=deps-osx-arm
	BUILD_KHOJI=$(DEPS_OSX_ARM) $(GOBUILD) -o $(BINARY_NAME) -v
endif

# all:
# 	@echo $(OS_ARCH)
# 	@echo $(GOPATH)
# 	@echo $(FINAL_CMD)

all: build
build: #$(BUILD_DEPS)
#	$(GITCMD) checkout $(CHECKOUT_BRANCH)
#	$(GOBUILD) -o $(BINARY_NAME) -v
	$(BUILD_KHOJI)
#	@echo $(BUILD_DEPS)
#	@echo $(BUILD_KHOJI)
#	@echo $(OS_ARCH)

clean: 
	$(GOCLEAN)
	rm -rf $(DIST_DIR)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_OSX)
	rm -f khoji.log
run:
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(GOBUILD) -o $(BINARY_NAME) -v 
	# ./$(BINARY_NAME) start
deps-linux:
	$(DEPS_LINUX) $(GOGET) -u github.com/satindergrewal/kmdgo
	$(DEPS_LINUX) $(GOGET) -u github.com/Meshbits/khoji
	$(DEPS_LINUX) $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	$(DEPS_LINUX) $(GOGET) -u github.com/fasthttp/router
	$(DEPS_LINUX) $(GOGET) -u github.com/valyala/fasthttp

deps-osx:
	$(DEPS_OSX) $(GOGET) -u github.com/satindergrewal/kmdgo
	$(DEPS_OSX) $(GOGET) -u github.com/Meshbits/khoji
	$(DEPS_OSX) $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	$(DEPS_OSX) $(GOGET) -u github.com/fasthttp/router
	$(DEPS_OSX) $(GOGET) -u github.com/valyala/fasthttp

deps-osx-arm:
	$(DEPS_OSX_ARM) $(GOGET) -u github.com/satindergrewal/kmdgo
	$(DEPS_OSX_ARM) $(GOGET) -u github.com/Meshbits/khoji
	$(DEPS_OSX_ARM) $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	$(DEPS_OSX_ARM) $(GOGET) -u github.com/fasthttp/router
	$(DEPS_OSX_ARM) $(GOGET) -u github.com/valyala/fasthttp

deps-win:
	$(DEPS_WIN) $(GOGET) -u github.com/satindergrewal/kmdgo
	$(DEPS_WIN) $(GOGET) -u github.com/Meshbits/khoji
	$(DEPS_WIN) $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	$(DEPS_WIN) $(GOGET) -u github.com/fasthttp/router
	$(DEPS_WIN) $(GOGET) -u github.com/valyala/fasthttp

# Cross compilation
build-linux: deps-linux
	rm -rf $(DIST_UNIX_PATH)
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_UNIX_PATH)
	$(DEPS_LINUX) $(GOBUILD) -o $(DIST_UNIX_PATH)/$(BINARY_NAME) -v
	$(CP_AV) $(DIST_FILES) $(DIST_UNIX_PATH)
	cd $(DIST_UNIX_PATH); zip -r ../$(BINARY_NAME)_linux.zip *; ls -lha ../; pwd
	$(RM_RFV) $(DIST_UNIX_PATH)
	cd $(ROOT_DIR)
build-osx: deps-osx
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_OSX_PATH)
	$(DEPS_OSX) $(GOBUILD) -o $(DIST_OSX_PATH)/$(BINARY_NAME) -v
	$(CP_AV) $(DIST_FILES) $(DIST_OSX_PATH)
	cd $(DIST_OSX_PATH); zip -r ../$(BINARY_NAME)_macos.zip *
	$(RM_RFV) $(DIST_OSX_PATH)
	cd $(ROOT_DIR)
build-osx-arm: deps-osx-arm
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_OSX_ARM_PATH)
	$(DEPS_OSX_ARM) $(GOBUILD) -o $(DIST_OSX_ARM_PATH)/$(BINARY_NAME) -v
	$(CP_AV) $(DIST_FILES) $(DIST_OSX_ARM_PATH)
	cd $(DIST_OSX_ARM_PATH); zip -r ../$(BINARY_NAME)_macos_arm64.zip *
	$(RM_RFV) $(DIST_OSX_ARM_PATH)
	cd $(ROOT_DIR)
build-win: deps-win
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_WIN_PATH)
	$(DEPS_WIN) $(GOBUILD) -o $(DIST_WIN_PATH)/$(BINARY_WIN) -v
	$(CP_AV) $(DIST_FILES) $(DIST_WIN_PATH)
	cd $(DIST_WIN_PATH); zip -r ../$(BINARY_NAME)_win.zip *
	$(RM_RFV) $(DIST_WIN_PATH)
	cd $(ROOT_DIR)