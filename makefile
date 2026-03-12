APP_NAME := wakey
CMD_PATH := ./cmd/wakey
BIN_DIR := ./bin
DIST_DIR := ./dist

VERSION ?= dev

.PHONY: build run clean release

build:
	@mkdir -p $(BIN_DIR)
	@go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

run: build
	@$(BIN_DIR)/$(APP_NAME)

clean:
	@rm -rf $(BIN_DIR) $(DIST_DIR)

release:
	@mkdir -p $(DIST_DIR)
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o $(DIST_DIR)/$(APP_NAME)-darwin-arm64 $(CMD_PATH)
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(DIST_DIR)/$(APP_NAME)-darwin-amd64 $(CMD_PATH)
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(DIST_DIR)/$(APP_NAME)-linux-amd64 $(CMD_PATH)
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(DIST_DIR)/$(APP_NAME)-windows-amd64.exe $(CMD_PATH)
