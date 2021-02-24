FROM golang:alpine AS builder

WORKDIR /go/src/github.com/tosone/golang-gin-template

ADD . .

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache gcc make musl-dev git tar

ARG swag_version=1.7.0

ADD https://github.com/swaggo/swag/releases/download/v${swag_version}/swag_${swag_version}_Linux_x86_64.tar.gz .

RUN tar -zxvf swag_${swag_version}_Linux_x86_64.tar.gz && mv swag /usr/bin && rm -rf swag_${swag_version}_Linux_x86_64.tar.gz

RUN go mod download && make release && cp bin/golang-gin-template /tmp

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
