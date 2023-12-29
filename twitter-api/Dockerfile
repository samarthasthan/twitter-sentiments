FROM golang:1.21.5

WORKDIR /usr/src/twitter-api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY tweets.csv .

COPY . .
RUN go build -v -o app .