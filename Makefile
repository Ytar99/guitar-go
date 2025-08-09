.PHONY: build run swagger

swagger:
    swag init -g ./cmd/main.go --output docs

build: swagger
    go build -o bin/main cmd/main.go

run: swagger
    go run cmd/main.go --config=./configs/config.yaml