FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy entire project including docs
COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
COPY configs/config.yaml ./configs/config.yaml

CMD ["./main", "--config=./configs/config.yaml"]