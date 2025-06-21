Go gRPC 使用示例

## 准备工作

1. 安装 protoc，参考 [Protocol Buffer Compiler Installation](https://protobuf.dev/installation/)

2. 安装 `protoc-gen-go` 和 `protoc-gen-go-grpc` 插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

3. 生成对应 gRPC 代码

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/greet.proto
```

4. 客户端和服务端通过生成的代码进行通信