FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/comment/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/usercomment /usr/local/go/src/douyin/pkg/comment/usercomment.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /app
COPY --from=builder /app/usercomment /app/usercomment
COPY --from=builder /app/etc /app/etc
CMD ["./usercomment", "-f", "etc/usercomment.yaml"]
