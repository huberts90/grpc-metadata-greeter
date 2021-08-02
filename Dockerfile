FROM golang:1.15.8-alpine
RUN apk add --no-cache --update \
    git \
    protobuf
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go get -u \
    github.com/golang/protobuf \
    google.golang.org/grpc \
    google.golang.org/protobuf/cmd/protoc-gen-go\
    google.golang.org/grpc/cmd/protoc-gen-go-grpc 2>&1

COPY . /usr/src/app
WORKDIR /usr/src/app
RUN scripts/compileproto.sh

RUN go test ./internal/...