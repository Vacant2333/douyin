FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY ../../../pkg/userinfo-demo/api/ /usr/local/go/src/douyin/pkg/userinfo-demo/api/
COPY ../../../pkg/userinfo-demo/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/userinfo /usr/local/go/src/douyin/pkg/userinfo-demo/api/template.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/userinfo /app/userinfo
COPY --from=builder /app/etc /app/etc

CMD ["./userinfo", "-f", "etc/template.yaml"]
