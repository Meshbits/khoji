# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
#GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
CGO_CFLAGS=$(shell env CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/")
CGO_LDFLAGS_DARWIN=$(shell env CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security")
CGO_LDFLAGS_WIN="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv"
CGO_LDFLAGS_LINUX="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm"
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
CHECKOUT_BRANCH=pbca26

# OS condition reference link: https://gist.github.com/sighingnow/deee806603ec9274fd47
UNAME_S=$(shell uname -s)

all:
	@echo $(OSFLAG)

all: build
build:
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(GOBUILD) -o $(BINARY_NAME) -v
# test: 
#	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(DIST_DIR)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_OSX)
run:
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(GOBUILD) -o $(BINARY_NAME) -v 
	# ./$(BINARY_NAME) start
deps-linux:
	GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" $(GOGET) -u github.com/satindergrewal/kmdgo
	GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" $(GOGET) -u github.com/Meshbits/khoji
	GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" $(GOGET) -u github.com/fasthttp/router
	GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" $(GOGET) -u github.com/valyala/fasthttp

deps-osx:
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" $(GOGET) -u github.com/satindergrewal/kmdgo
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" $(GOGET) -u github.com/Meshbits/khoji
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" $(GOGET) -u github.com/fasthttp/router
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" $(GOGET) -u github.com/valyala/fasthttp

deps-osx-arm:
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" $(GOGET) -u github.com/satindergrewal/kmdgo
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" $(GOGET) -u github.com/Meshbits/khoji
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" $(GOGET) -u github.com/fasthttp/router
	GO111MODULE=auto CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" $(GOGET) -u github.com/valyala/fasthttp

deps-win:
	GO111MODULE=auto CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" $(GOGET) -u github.com/satindergrewal/kmdgo
	GO111MODULE=auto CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" $(GOGET) -u github.com/Meshbits/khoji
	GO111MODULE=auto CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" $(GOGET) -u gopkg.in/rethinkdb/rethinkdb-go.v6
	GO111MODULE=auto CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" $(GOGET) -u github.com/fasthttp/router
	GO111MODULE=auto CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" $(GOGET) -u github.com/valyala/fasthttp

# Cross compilation
build-linux: deps-linux
	rm -rf $(DIST_UNIX_PATH)
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_UNIX_PATH)
	GO111MODULE=auto CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(DIST_UNIX_PATH)/$(BINARY_NAME) -v
	$(CP_AV) $(DIST_FILES) $(DIST_UNIX_PATH)
	cd $(DIST_UNIX_PATH); zip -r ../$(BINARY_NAME)_linux.zip *; ls -lha ../; pwd
	$(RM_RFV) $(DIST_UNIX_PATH)
	cd $(ROOT_DIR)
build-osx: deps-osx
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_OSX_PATH)
	GO111MODULE=auto CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(DIST_OSX_PATH)/$(BINARY_NAME) -v
	$(CP_AV) $(DIST_FILES) $(DIST_OSX_PATH)
	cd $(DIST_OSX_PATH); zip -r ../$(BINARY_NAME)_macos.zip *
	$(RM_RFV) $(DIST_OSX_PATH)
	cd $(ROOT_DIR)
build-osx-arm: 
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_OSX_ARM_PATH)
	GO111MODULE=auto CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/darwin_arm64 -lsaplinglib -framework Security" CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(DIST_OSX_ARM_PATH)/$(BINARY_NAME) -v
	$(CP_AV) $(DIST_FILES) $(DIST_OSX_ARM_PATH)
	cd $(DIST_OSX_ARM_PATH); zip -r ../$(BINARY_NAME)_macos_arm64.zip *
	$(RM_RFV) $(DIST_OSX_ARM_PATH)
	cd $(ROOT_DIR)
build-win:
	$(GITCMD) checkout $(CHECKOUT_BRANCH)
	$(MKDIR_P) $(DIST_WIN_PATH)
	GO111MODULE=auto CGO_CFLAGS="-I$(HOME)/go/src/github.com/satindergrewal/saplinglib/src/" CGO_LDFLAGS="-L$(HOME)/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv" CC="x86_64-w64-mingw32-gcc" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(DIST_WIN_PATH)/$(BINARY_WIN) -v
	$(CP_AV) $(DIST_FILES) $(DIST_WIN_PATH)
	cd $(DIST_WIN_PATH); zip -r ../$(BINARY_NAME)_win.zip *
	$(RM_RFV) $(DIST_WIN_PATH)
	cd $(ROOT_DIR)