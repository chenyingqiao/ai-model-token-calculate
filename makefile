# Makefile for ai-model-token-calculate CLI tool

# 可执行文件名称
BINARY=tokencli

# 安装目录，优先使用 GOPATH/bin，否则 /usr/local/bin
INSTALL_DIR?=$(shell go env GOPATH)/bin
ifeq ($(INSTALL_DIR),)
INSTALL_DIR=/usr/local/bin
endif

.PHONY: all build install clean

all: build

build:
	go build -o $(BINARY) ./main.go

install: build
	@echo "Installing $(BINARY) to $(INSTALL_DIR)..."
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY) $(INSTALL_DIR)
	@echo "Installed $(BINARY) to $(INSTALL_DIR)"

clean:
	rm -f $(BINARY)

