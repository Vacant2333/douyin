# 多阶段构建,先用golang环境编译好main
FROM golang:latest AS builder
# 工作目录
WORKDIR /demo
# 将demo用到的内部库放入容器内
ADD pkg/demo /usr/local/go/src/douyin/pkg/demo
# 将demo项目所需要的内容放入docker内
ADD go.mod .
ADD go.sum .
ADD cmd/demo .
# 设置环境变量
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 编译
RUN CGO_ENABLED=0 go build main.go
# 多阶段构建,把从builder编译好的main文件copy过来
FROM gcr.io/distroless/static-debian11:latest
WORKDIR /demo
# copy编译好的main可执行文件
COPY --from=builder /demo/main main
# 运行
CMD ["./main"]