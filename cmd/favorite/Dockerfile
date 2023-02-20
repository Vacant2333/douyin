FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg/favorite /usr/local/go/src/douyin/pkg/favorite
COPY common /usr/local/go/src/douyin/common
COPY pkg/favorite/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/favorite /usr/local/go/src/douyin/pkg/favorite/userOptService.go

FROM gcr.io/distroless/static-debian11:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/favorite /app/favorite
COPY --from=builder /app/etc /app/etc
CMD ["./favorite", "-f", "etc/userOptService.yaml"]
