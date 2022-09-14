# gRPC

Exercitando modos de comunicação com gRPC

links:
    - https://grpc.io/
    - https://developers.google.com/protocol-buffers/docs/overview

# Anotações de trabalho

WORKDIR: go/src

Create image go lang + grpc
```
docker build -t brunocicom/fc-grpc . grpc
```

Container de desenvolvimento
``` bash
docker run --rm -it --name grpc -v $(pwd)/:/go/src brunocicom/fc-grpc sh
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