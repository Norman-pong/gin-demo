# 定义变量
APP_NAME := gin-demo
GO_API := ./cmd/api

# 默认目标
.PHONY: all
all: build

# 构建命令
.PHONY: build
build:
	go build -o out/ ${GO_API}/main.go

# 运行命令
.PHONY: run
run:
	go run $(GO_API)/main.go