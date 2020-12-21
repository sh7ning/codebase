# Gin Application

## 配置

1. `cp app/api/app/cmd/config.example.yaml config.yaml`
2. `vim config.yaml`

## Development

* Run

```
go run app/api/app/cmd/main.go start -c config.yaml
```

or

```
./auto-build-dev
```

## Production

* Run

```
./app start -c config.yaml
```

# docker tag

* api-app 项目
    
```
docker build --build-arg APP_MODULE=api --build-arg APP_NAME=app -t docker.example.com/sh7ning/codebase-api-app:v0.0.1 .

docker push docker.example.com/sh7ning/codebase-api-app:v0.0.1
```

> 测试运行: `docker run --rm -it -p 8080:8080 -v $(pwd):/app docker.example.com/sh7ning/codebase-api-app:v0.0.1 start -c config.yaml`

## Todo

