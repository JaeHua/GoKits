
## 项目介绍
  这是一个goweb开发的入门级别项目的后端工程,也是一个简单的注册登陆的通用后台入门模板。利用Gin框架进行开发。在这个项目中会了解到如何进行token分发 、连接数据库 、中间件 、路由 、gorm、jwt 、viper 以及postman使用和类企业级项目重构等等知识点。

## 项目启动

**第一种：**

```go build -o myApp main.go route.go ```
然后在当前路径运行`.exe`文件 **./myApp** .

**第二种:**

```go run main.go route.go```

## 项目预备知识
● Go初步语法

● Gin框架

● Jwt

● Viper

● Gorm


## 涉及知识简单介绍

### Jwt

用于Web token

官网:https://jwt.p2hp.com/introduction

JWT代码仓库:https://github.com/dgrijalva/jwt-go

拉取:

```go get github.com/dgrijalva/jwt-go```


### Viper
用于配置文件的读取管理等

代码仓库:https://github.com/spf13/viper

拉取:

```go get github.com/spf13/viper```

### Gorm
把struct类型和数据库表记录进行映射，操作数据库的时候不需要直接手写Sql代码

代码仓库:https://github.com/go-gorm/gorm

拉取:

```go get -u github.com/jinzhu/gorm```

### Gin

流行的goweb框架

代码仓库:https://github.com/gin-gonic/gin

拉取:

```go get -u github.com/gin-gonic/gin```

### Mysql驱动

go中mysql驱动

代码仓库:https://github.com/go-sql-driver/mysql

拉取:

```go get github.com/go-sql-driver/mysql```

### Postman

接口测试工具

官网:https://www.postman.com/

## 参考:

[https://www.bilibili.com/video/BV1Rg411u7xH/](https://www.bilibili.com/video/BV1CE411H7bQ/?spm_id_from=333.337.search-card.all.click)
