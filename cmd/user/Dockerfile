FROM golang:latest AS builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY pkg /usr/local/go/src/douyin/pkg
COPY common /usr/local/go/src/douyin/common
COPY pkg/user/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/user /usr/local/go/src/douyin/pkg/user/user.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc
CMD ["./user", "-f", "etc/userService.yaml"]
