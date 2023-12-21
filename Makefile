.PHONY: clean build build-in-docker

APP_NAME = ftpdadmin
BUILD_DIR = $(PWD)/build
DIST_DIR = $(PWD)/dist


run:
	#go run main.go
	@air

all: clean build-dev

clean:
	rm -rf ./build


build-dev:
	go mod tidy
	rm -f $(APP_NAME)
	@go build -o $(APP_NAME) main.go

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) main.go
	ls -sh $(BUILD_DIR)/$(APP_NAME)
	upx --best --lzma $(BUILD_DIR)/$(APP_NAME)
	ls -sh $(BUILD_DIR)/$(APP_NAME)

build-dist:
	@rm -rf ./dist
	CGO_ENABLED=0 go build -ldflags="-s -w" -o $(DIST_DIR)/$(APP_NAME) main.go
	ls -sh $(DIST_DIR)/$(APP_NAME)
	upx --best --lzma $(DIST_DIR)/$(APP_NAME)
	ls -sh $(DIST_DIR)/$(APP_NAME)
	cp -aR views $(DIST_DIR)/views
	cp -aR public $(DIST_DIR)/public
	cp config.yml $(DIST_DIR)/config.yml

build-win:
	GOOS=windows go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-win64.exe
	ls -sh $(BUILD_DIR)/$(APP_NAME)-win64.exe
	upx --best --lzma $(BUILD_DIR)/$(APP_NAME)-win64.exe
	ls -sh $(BUILD_DIR)/$(APP_NAME)-win64.exe

build-mac:
	GOOS=darwin go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-mac64
	ls -sh $(BUILD_DIR)/$(APP_NAME)-mac64
	upx --best --lzma $(BUILD_DIR)/$(APP_NAME)-mac64
	ls -sh $(BUILD_DIR)/$(APP_NAME)-mac64

build-in-docker:
	apt-get update && apt-get install -y --no-install-recommends xz-utils
	curl -L# https://github.com/upx/upx/releases/download/v4.2.1/upx-4.2.1-amd64_linux.tar.xz | tar -xJv
	mv upx-*-amd64_linux/upx /usr/local/bin/ && rm -rf upx-*-amd64_linux
	make build-dist
	mv /dist dist && chown -R 1000:1000 dist
