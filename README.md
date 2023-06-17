# gRPC Server Sample

サイコロの値を返すgRPCサーバーのサンプル

[参考サイト](https://zenn.dev/hsaki/books/golang-grpc-starting)

## 開発環境の構築

1. Protocol Bufferをインストール
   [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)

2. protoc-gen-go-grpc をインストール

```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## コード生成

```shell
protoc --go_out=. --go-grpc_out=. proto/*
```

## クライアントのサンプル

[https://github.com/takara2314/grpc-client-sample](https://github.com/takara2314/grpc-client-sample)
