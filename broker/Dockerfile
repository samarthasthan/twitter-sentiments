FROM golang:1.21.5

WORKDIR /usr/src/broker

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app .