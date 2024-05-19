# GoKits

GoKits是一个基于Golang和Vue的全栈项目模板,前端采用Vue3+Tailwindcss+ElementPlus开发,后端使用Gin框架开发。整个项目包含前后端代码,抽取出通用的功能模块和组件,并通过Docker实现持续集成和部署。


![jojocat](https://github.com/JaeHua/GoKits/assets/126366914/a86ab124-c10e-4588-b5c9-c5bc6e7abfe2)




## 项目结构

```
go-kits
├─backend
│  ├─common
│  ├─config
│  ├─controller
│  ├─dto
│  ├─middleware
│  ├─model
│  ├─response
│  └─util
└─frontend
│    ├─public
│    └─src
│        ├─assets
│        ├─components
│        │  └─icons
│        ├─router
│        ├─service
│        ├─store
│        │  └─module
│        ├─utils
│        └─views
│            └─layout
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

克隆项目,配置文件然后启动docker容器
### Mysql



- 拉取镜像（版本5.7）

```Bash
docker pull mysql:5.7
```

- 创建对应文件夹

```Bash
mkdir -p /home/mysql/conf
mkdir -p /home/mysql/data
```

- 授权文件夹

```Bash
chmod 777 /home/mysql/conf
chmod 777 /home/mysql/data
```

- 运行容器

```Bash
docker run -d -p 3310:3306 -v /home/mysql/conf:/etc/mysql/conf.d/ -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql_gokits
```
## Redis部署

- 拉取镜像(版本7.0)

```Bash
docker pull redis:7.0
```

- 创建对应文件夹

```Bash
mkdir -p /usr/local/docker/redis/conf
mkdir -p /usr/local/docker/redis/data
```

- 授权文件夹

```Bash
chmod 777 /usr/local/docker/redis/conf
chmod 777 /usr/local/docker/redis/data
```

- 运行容器

```Bash
docker run \
    -d \
    -p 6379:6379 \
    --name redis_gokits \
    --restart=always \
    -v /usr/local/docker/redis/data:/data \
    -v /usr/local/docker/redis/conf/redis.conf:/etc/redis \
    redis
```

## Vue部署

- 打包成dist

```Bash
npm build
```

- 创建对应文件夹

```Bash
mkdir /usr/local/dockernginxdist
```

- 授权文件夹

```Bash
chmod 777 /usr/local/dockernginxdist
```

- 构建镜像

```Bash
docker build -t vue_gokits .
```

- 运行容器

```Bash
docker run -p 3300:3300 -d vue_gokits
```

## Gin部署

- 创建对应文件夹(代码打包放在这)

```Bash
mkdir /home/gokits/
```

- 构建镜像

```Bash
docker build -t gokits:1.0 .
```

- 运行容器

```Bash
docker run -d -p 3344:3344 --name gin_gokits gokits:1.0
```

## 容器监测

利用**portainer**这个工具监控我们docker的运行状况

![image](https://github.com/JaeHua/GoKits/assets/126366914/1bcaaa56-e4a1-42a6-bd45-a2984ac7a7d5)


## Redis管理

推荐使用**RedisDesktopManager** ,开源免费的轻巧工具
![image](https://github.com/JaeHua/GoKits/assets/126366914/94845abd-b4a6-46b4-b0a4-c6fc062ca00a)

下载链接：https://github.com/RedisInsight/RedisDesktopManager/releases

## Mysql管理

推荐使用**DBeaver**，开源免费工具

