FROM golang:1.23.5 AS builder

WORKDIR /app

COPY main.go .

RUN go mod init server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:latest

COPY --from=builder /app/server /usr/local/bin/server

EXPOSE 8080 8081

CMD ["server"]
