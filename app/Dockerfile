# build
FROM golang:1.18.3-alpine3.16  AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build cmd/main.go 

# deploy
FROM alpine:latest

WORKDIR /

COPY --from=build /app/main /main

ENTRYPOINT ["main"]
