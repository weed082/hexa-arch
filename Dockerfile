FROM golang:1.17.6-alpine3.15

WORKDIR /app
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
RUN go get github.com/cosmtrek/air
COPY . . 
CMD ["air", "-c", "air.toml"]
