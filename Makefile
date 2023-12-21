.PHONY: clean build

APP_NAME = ftpdadmin
BUILD_DIR = $(PWD)/build
DIST_DIR = $(PWD)/dist


run:
	@go run main.go

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
	cp app-config.yml $(DIST_DIR)/app-config.yml

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

