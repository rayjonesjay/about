FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod .

COPY . .

RUN GOARCH=amd64 go build -o main .

EXPOSE 9090

CMD ["/app/main"]
