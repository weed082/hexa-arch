# build
FROM golang:1.17.6-alpine3.15 AS build

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

ENTRYPOINT ["/main"]