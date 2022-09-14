FROM golang:1.19.1-alpine

WORKDIR /go/src

RUN apk update && apk add protobuf & \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN export PATH="$PATH:$(go env GOPATH)/bin"

CMD ["tail","-f","/dev/null"]