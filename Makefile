.PHONY: build clean test coverage lint help

APP_NAME=gos
MAIN_PATH=main.go

help: ## 显示帮助信息
	@echo "使用方法:"
	@echo "make <target>"
	@echo ""
	@echo "可用的目标:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## 构建应用
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

run: ## 运行应用
	go run $(MAIN_PATH)

clean: ## 清理构建产物
	rm -rf bin/
	rm -rf dist/

lint: ## 运行代码检查
	golangci-lint run

mod: ## 整理并下载依赖
	go mod tidy
	go mod download

.DEFAULT_GOAL := help
