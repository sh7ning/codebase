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

## Todo
