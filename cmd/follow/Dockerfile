FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/follow/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/follow /usr/local/go/src/douyin/pkg/follow/follow.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /app
COPY --from=builder /app/follow /app/follow
COPY --from=builder /app/etc /app/etc
CMD ["./follow", "-f", "etc/follow.yaml"]
