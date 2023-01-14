all: build

build:
	@go build -o build/kklogTUI cmd/main.go

run:
	@go run cmd/main.go

clean:
	@rm build/kklogTUI

install:
	@echo "将 kklogTUI 复制到 /usr/local/bin 下"
	@sudo mv ./build/kklogTUI /usr/local/bin/

.PHONY: all build run clean install
