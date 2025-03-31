FROM golang:1.23 AS builder

WORKDIR /gravtest

COPY ./go.mod ./go.sum ./


RUN go mod download


RUN go mod tidy

# Copy the source code
COPY ./ ./



RUN go build cmd/main.go

EXPOSE 8080

CMD ["./main"]
