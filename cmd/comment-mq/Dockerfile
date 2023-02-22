FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/comment-mq/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/comment-mq /usr/local/go/src/douyin/pkg/comment-mq/mq.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /app
COPY --from=builder /app/comment-mq /app/comment-mq
COPY --from=builder /app/etc /app/etc
CMD ["./comment-mq", "-f", "etc/comment.yaml"]
