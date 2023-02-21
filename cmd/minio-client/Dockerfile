FROM golang:latest AS builder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/minio-client/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/minioclient /usr/local/go/src/douyin/pkg/minio-client/minioclient.go

FROM gcr.io/distroless/static-debian11:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/minioclient /app/minioclient
COPY --from=builder /app/etc /app/etc
CMD ["./minioclient", "-f", "etc/minioclient.yaml"]
