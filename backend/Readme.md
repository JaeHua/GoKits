# Gokits后端

## 项目介绍

这个项目是使用 Golang Gin 框架开发的一个简单的后台系统功能,用于 Golang Web 开发入门。

## 功能简介

- 用户注册与登录功能
- JWT 加密 token 进行用户鉴权
- Gorm 操作 MySQL 数据库
- 中间件处理请求
- 路由功能设计
- Viper 读取配置文件
- 持久化缓存
- 错误处理与日志记录
- Postman 接口测试
- 项目结构与工程化设计

## 使用技术

- Golang 1.17
- Gin Web 框架
- Gorm ORM
- JWT 加密
- Viper 配置
- MySQL 数据库
- Redis 缓存
- Logger 日志
- validator 验证
- godoc 文档
- errgroup 并发
- go.mod 依赖管理


## 目录结构

```
src
├── common       
├── config
├── controller     
├── middleware
├── dto  
├── routers   
├── response
├── util
├── route.go
└── main.go
```

## 运行示例

* 克隆项目后,自行配置`.env`文件,运行:

```
go build
./your_app_name
```
* 利用docker运行


```
docker build -t gokits:1.0
docker run -d -p port1:port2 --name gin_gokits gokits:1.0
```

访问:http://localhost:port

## 使用备忘

本项目是个人学习使用的简单 Golang Web 项目模板,目的是快速上手 Golang Web 开发。
项目结构及代码设计上还有许多不足,欢迎大家提出宝贵意见。
