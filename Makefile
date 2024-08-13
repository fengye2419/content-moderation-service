# Makefile

# 二进制文件名称
BINARY_NAME=content-moderation-service

# Go 源文件目录
SRC_DIR=cmd

# Swagger 文档目录
SWAGGER_DIR=docs

# Go 命令
GO=go

# 单元测试目录
GO_TEST_PACKAGES=$(shell $(GO) list ./internal/api ./internal/service ./internal/model) 

# 输出帮助信息
.PHONY: help
help:
	@echo "可用命令："
	@echo "  build   - 编译二进制文件"
	@echo "  swagger - 生成 Swagger 文档"
	@echo "  test    - 运行单元测试"
	@echo "  clean   - 清理生成的文件"
	@echo "  help    - 输出帮助信息"

# 默认目标
.PHONY: all
all: build

# 编译二进制文件
.PHONY: build
build:
	$(GO) build -o $(BINARY_NAME) $(SRC_DIR)/main.go

# 生成 Swagger 文档
.PHONY: swagger
swagger:
	swag init -g $(SRC_DIR)/main.go -o $(SWAGGER_DIR)

# 运行单元测试
.PHONY: test
test:
	$(GO) test $(GO_TEST_PACKAGES)

# 清理生成的文件
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -rf $(SWAGGER_DIR)/*.json $(SWAGGER_DIR)/*.yaml

