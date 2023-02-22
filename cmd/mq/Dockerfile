FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/mq/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/mq /usr/local/go/src/douyin/pkg/mq/mq.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /app
COPY --from=builder /app/mq /app/mq
COPY --from=builder /app/etc /app/etc
CMD ["./mq", "-f", "etc/mq.yaml"]
