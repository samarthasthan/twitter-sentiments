FROM golang:1.21.5

WORKDIR /usr/src/analyser

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app .