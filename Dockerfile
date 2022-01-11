FROM golang:latest

WORKDIR /app
COPY go.mod .
COPY go.SUM .
RUN go mod download
COPY . .
RUN go build cmd/main.go

CMD ["./main"]