FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod .

COPY . .

RUN GOARCH=amd64 go build -o main .

FROM alpine:latest

COPY --from=builder /app/main /main

EXPOSE 9000

CMD ["/main"]
