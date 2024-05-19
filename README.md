# GoKits

GoKits是一个基于Golang和Vue的全栈项目模板,前端采用Vue3+Tailwindcss+ElementPlus开发,后端使用Gin框架开发。整个项目包含前后端代码,抽取出通用的功能模块和组件,并通过Docker实现持续集成和部署。

## 项目结构

```
go-kits
├── backend  // Golang后端代码
│   └── main.go
│   └── route.go
├── frontend // Vue前端代码
│   ├── public
|   └── src
├── .dockerignore
└── README.md
```

## 后端技术

- Golang 
- Gin Web 框架
- Gorm ORM
- JWT 鉴权
- Mysql 数据库
- Redis 缓存
- Logger 日志
- Validator 验证
- so on
## 前端技术 

- Vue3
- Vite
- Tailwind CSS
- Element Plus
- Axios 网络请求
- Pinia 状态管理
- Vue Router 路由
- so on

## 部署方式

- Docker 镜像化
- Nginx 代理
- Linux 发布正式版

## 使用说明

克隆项目,配置文件,在 IDE 运行`docker-compose up`启动开发环境。

## 项目亮点

- 前后端分离架构
- Docker 一键运行
- 自动化测试与持续集成
- 业务通用抽象
- 好的代码结构与组织

期待您的Star!
