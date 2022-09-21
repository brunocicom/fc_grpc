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
# Client Universal :: Evans

https://github.com/ktr0731/evans

## Install

$ curl -L 'https://github.com/ktr0731/evans/releases/download/v0.10.9/evans_linux_amd64.tar.gz' | tar xvzf -
$ file evans
evans: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, not stripped
$ mv evans /usr/local/bin/evans # This path must be included in $PATH

## Running cliente

```
evans -r repl --host localhost --port 50051
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