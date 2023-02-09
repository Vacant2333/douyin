# go-zero-demo

```bash
cd rpc
goctl rpc protoc ./proto/userinfo.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
cd ../api
goctl api go -api userinfo.api -dir .
```