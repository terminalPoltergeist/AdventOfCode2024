BINARY_NAME=AOC24
WD=$(shell pwd)

all:
	@echo "Resolving dependencies..."
	@go mod tidy
	@echo "Compiling the binary..."
	@go build -C cmd/ -o ${WD}/${BINARY_NAME}
.PHONY: all

run:
	@echo "Resolving dependencies..."
	@go mod tidy
	@echo "Compiling the binary..."
	@go build -C cmd/ -o ${WD}/${BINARY_NAME}
	@${WD}/${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
.PHONY: clean
