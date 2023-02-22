FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/gateway/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/gateway /usr/local/go/src/douyin/pkg/gateway/api.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /app
COPY --from=builder /app/gateway /app/gateway
COPY --from=builder /app/etc /app/etc
CMD ["./gateway", "-f", "etc/user-api.yaml"]
