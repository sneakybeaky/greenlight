FROM golang:1.17-alpine3.13
WORKDIR /go-greenlight

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY cmd cmd
    RUN go build -o build/api github.com/sneakybeaky/greenlight/cmd/api
    SAVE ARTIFACT build/api /go-greenlight AS LOCAL build/api

unit-test:
    FROM +deps
    COPY cmd cmd
    RUN CGO_ENABLED=0 go test -v github.com/sneakybeaky/greenlight/...

all:
  BUILD +unit-test
  BUILD +build
