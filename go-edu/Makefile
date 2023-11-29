.PHONY: all build run gotool  help

BINARY="gin-admin-system"

all: gotool build

build:
ifeq ($(OS),Windows_NT)
	set CGO_ENABLED=0
	set GOARCH=amd64
	set GOOS=linux
	go build -o $(BINARY)
else
	CGO_ENABLED=0
	GOARCH=amd64
	GOOS=linux
	go build -o $(BINARY)
endif

run:
	@go run ./main.go

gotool:
	go fmt ./
	go vet ./

help:
	@echo "make - Format Go code and compile to generate binary files"
	@echo "make build - Compile the Go code to generate the binary file"
	@echo "make run - Run the Go code directly"
	@echo "make gotool - Formatted code"


#ifeq ($(OS),Windows_NT)
#	@echo "windows"
#else
#	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
#	@echo "make build - 编译 Go 代码, 生成二进制文件"
#	@echo "make run - 直接运行 Go 代码"
#	@echo "make clean - 移除二进制文件和 vim swap files"
#	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
#endif











