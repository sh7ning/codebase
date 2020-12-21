## 1. 构建应用 ##
FROM golang:1.13-stretch as builder

ARG APP_MODULE
ARG APP_NAME

RUN export GO111MODULE=on \
    && mkdir -p /go/src/${APP_MODULE}

COPY . /go/src/${APP_MODULE}

#RUN echo "docker build info: /go/src/${APP_MODULE} /go/bin/${APP_MODULE}_${APP_NAME} ./app/${APP_MODULE}/${APP_NAME}/cmd/main.go"

RUN cd /go/src/${APP_MODULE} \
    && CGO_ENABLED=1 go build -ldflags '-s -w' -o /go/bin/app ./app/${APP_MODULE}/${APP_NAME}/cmd/main.go

## 2. 应用 ##
FROM debian:stretch

ARG APP_MODULE
ARG APP_NAME

# 增加根证书
RUN apt-get update \
    && apt-get install ca-certificates -y

COPY --from=builder /go/bin/app /usr/local/bin/

# .env => /app/.env
WORKDIR /app
VOLUME /app

# 运行时指定,此处不进行设置
#EXPOSE 8900 8901

CMD ["app"]

ENTRYPOINT ["app"]
