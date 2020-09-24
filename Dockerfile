FROM golang:1.14-alpine as builder
ARG CHINESE_ENABLE
ARG GOPROXY

ENV GOPROXY ${GOPROXY}
ENV CHINESE_ENABLE true

WORKDIR /app

RUN if [ ${CHINESE_ENABLE} ]; then \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
;fi

COPY . /app
RUN export GOPROXY="https://goproxy.cn,direct" && go mod download && apk add make && make clean && make && cp server.example.json server.json

FROM alpine:latest as prod

RUN if [ ${CHINESE_ENABLE} ]; then \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
;fi

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=0 /app/build/linux/redis .
COPY --from=0 /app/server.json .

CMD ["/app/redis"]