FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD ../../../go.mod .
ADD ../../../go.sum .
RUN go mod download
# 构建镜像报错 package pkg/comment/rpc/usercomment.go is not in GOROOT (/usr/local/go/src/pkg/comment/rpc/usercomment.go)
# 解决 部分文件路上上添加usr/local/go/src/douyin 如下
COPY /pkg/comment/api /usr/local/go/src/douyin/pkg/comment/api
COPY /pkg/comment/common /usr/local/go/src/douyin/pkg/comment/common
COPY /pkg/comment/rpc /usr/local/go/src/douyin/pkg/comment/rpc
COPY /pkg/comment/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/usercomment /usr/local/go/src/douyin/pkg/comment/rpc/usercomment.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/usercomment /app/usercomment
COPY --from=builder /app/etc /app/etc

CMD ["./usercomment", "-f", "etc/usercomment.yaml"]
