# Job Search Project

## 项目简介
这是一个用于求职相关功能的项目，提供了一系列与求职相关的服务和接口。

## 项目结构
```
├── cmd/
│   └── server/
│       └── server.go
├── config.yaml
├── ent/
├── go.mod
├── go.sum
└── internal/
    ├── application/
    ├── domain/
    ├── infrastructure/
    └── interface/
```

## 安装依赖
```bash
go mod tidy
go install github.com/golang/mock/mockgen@latest
```

## 运行项目
```bash
# 请确保在项目根目录下执行
go run cmd/server/server.go
```