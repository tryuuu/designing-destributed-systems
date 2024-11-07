FROM golang:1.23-alpine AS builder
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o monitor-mysql

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/monitor-mysql .

EXPOSE 8080

CMD ["./monitor-mysql", "-query", "SELECT 1", "-user", "root", "-password", "password", "-database", "test-db", "-address", "0.0.0.0:8080"]
