# Variables
APP_NAME=jarvis

run:
	@go run cmd/jarvis/main.go

build:
	@go build -o $(APP_NAME) cmd/jarvis/main.go