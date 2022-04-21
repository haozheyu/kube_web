FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o kube_web_api .

FROM debian:stretch-slim
COPY ./conf /conf
COPY --from=builder /build/kube_web_api /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends 

# 需要运行的命令
ENTRYPOINT ["/kube_web_api"]
