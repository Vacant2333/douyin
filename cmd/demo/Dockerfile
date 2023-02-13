FROM golang:latest AS builder
WORKDIR /demo
ADD pkg/demo /usr/local/go/src/douyin/pkg/demo
ADD go.mod .
ADD go.sum .
ADD cmd/demo .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 go build main.go

FROM gcr.io/distroless/static-debian11:latest
WORKDIR /demo
COPY --from=builder /demo/main main
CMD ["./main"]