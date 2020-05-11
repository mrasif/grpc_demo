# gRPC Demo

Go microservice demo with gRPC

## Generate proto file

```sh
protoc --go_out=plugins=grpc:chat chat.proto
```

Installation guide of `protoc` is [http://google.github.io/proto-lens/installing-protoc.html](http://google.github.io/proto-lens/installing-protoc.html)

## Run

```sh
go run main.go
go run chat.go
```
