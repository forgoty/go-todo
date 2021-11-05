FROM golang:1.16

WORKDIR /go/src/github.com/forgoty/go-todo

COPY go.mod go.sum ./

RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

RUN go mod download
