# grpc-go-101

## Setup

1. Install protoc command

```sh
# Protoc
brew install protobuf
# Clang-Format for format .proto file on vs-code
brew install clang-format
```

~~2. Install go plugin `go get -u github.com/golang/protobuf/protoc-gen-go`~~

2. Install

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

[REF](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code)

## Compile proto file

1. Write your proto
~~2. compile proto file `protoc --go_out=. *.proto`~~

2. Compile proto file

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pingpong.proto
```

## Ref

- https://github.com/varshard/helloproto
- https://medium.com/@titlebhoomtawathplinsut/%E0%B8%A1%E0%B8%B2%E0%B8%97%E0%B8%B3-grpc-service-%E0%B8%94%E0%B9%89%E0%B8%A7%E0%B8%A2-go-%E0%B8%81%E0%B8%B1%E0%B8%99-866d7452f5dd
- [gRPC new](https://programmingpercy.tech/blog/using-grpc-tls-go-react-no-reverse-proxy/)
- [Example gRPC Thai](https://thanwa.medium.com/grpc-%E0%B8%A1%E0%B8%B2%E0%B8%9D%E0%B8%B6%E0%B8%81%E0%B9%80%E0%B8%82%E0%B8%B5%E0%B8%A2%E0%B8%99-grpc-%E0%B8%81%E0%B8%B1%E0%B8%99-hands-on-e4cc325033fb)
