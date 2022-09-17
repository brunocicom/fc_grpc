# gRPC

Exercitando modos de comunicação com gRPC

links:
    - https://grpc.io/
    - https://developers.google.com/protocol-buffers/docs/overview

# Server

Run Server
```
go run cmd/server/server.go
```

# Anotações de trabalho

WORKDIR: go/src

Create image go lang + grpc
```
docker build -t brunocicom/fc-grpc . grpc
```

Container de desenvolvimento com usuário
``` bash
docker run --rm -it --name grpc --user 1000:1000 -v $(pwd)/:/go/src brunocicom/fc-grpc sh
```

build package pb
```
protoc --proto_path=proto proto/*.proto --go_out=pb
```

build package pb with go stub
```
protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
```

Run
``` bash
go run .
```

Build
``` bash
env GOOS=linux GOARCH=386 go build .
```

Image golang: https://hub.docker.com/_/golang