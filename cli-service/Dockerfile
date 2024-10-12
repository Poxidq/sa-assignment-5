FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cli-service ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cli-service /app/cli-service

EXPOSE 8080

ENTRYPOINT ["/app/cli-service"]

