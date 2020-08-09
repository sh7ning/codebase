## 1. 构建应用 ##
FROM golang:1.13-stretch as builder

ARG APP_MODULE
ARG APP_NAME

RUN export GO111MODULE=on \
    && mkdir -p /go/src/${APP_MODULE}

COPY . /go/src/${APP_MODULE}

RUN echo /go/src/${APP_MODULE}
RUN echo /go/bin/${APP_NAME} ${APP_NAME}.go

RUN cd /go/src/${APP_MODULE} \
    && CGO_ENABLED=1 go build -ldflags '-s -w' -o /go/bin/${APP_NAME} ./app/api/app/cmd/${APP_NAME}.go

## 2. 应用 ##
FROM debian:stretch

ARG APP_NAME

# 增加根证书
RUN apt-get update \
    && apt-get install ca-certificates -y

COPY --from=builder /go/bin/${APP_NAME} /usr/local/bin/

# .env => /app/.env
WORKDIR /app
VOLUME /app

# 运行时指定,此处不进行设置
#EXPOSE 8900 8901

CMD ["${APP_NAME}"]
