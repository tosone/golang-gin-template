FROM golang:1.16-alpine3.13 AS builder

WORKDIR /go/src/github.com/tosone/golang-gin-template

ADD . .

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache build-base musl-dev git tar

ARG swag_version=1.7.0

ADD https://github.com/swaggo/swag/releases/download/v${swag_version}/swag_${swag_version}_Linux_x86_64.tar.gz .

RUN tar -zxvf swag_${swag_version}_Linux_x86_64.tar.gz && mv swag /usr/local/bin && \
  rm -rf swag_${swag_version}_Linux_x86_64.tar.gz

# ENV GOPROXY="https://goproxy.cn,direct"
RUN go mod download && make build && cp bin/golang-gin-template /tmp

FROM alpine:3.13

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates bash && \
  mkdir -p /etc/config && \
  mkdir -p /data

COPY --from=builder /tmp/golang-gin-template /usr/bin
COPY config.yml.sample /etc/config/config.yml

VOLUME /etc/config

EXPOSE 4000

CMD ["golang-gin-template", "server"]
