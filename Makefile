APP=snake
APP_EXECUTABLE="./out/$(APP)"


install:
	@go install ./...
	@echo ""	
	@echo "Run: make help 		to get more info on the available zc_core make commands"
	@echo ""

run:
	@go run main.go

build:
	@go build -o bin/main main.go

install_lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

lint:
	@golangci-lint run

fresh:
	@go install github.com/mitranim/gow@latest
	@gow run .
	
