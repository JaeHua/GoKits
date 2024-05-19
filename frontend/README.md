# Vue3 + Tailwind CSS + Element Plus GoKits

## 项目运行

- 开发环境:

  ```
  npm install
  npm run dev
  ```

- 生产环境:

  ```
  npm build
  ```

## Vue部署

- 构建dist文件夹

  ```
  npm build
  ```

- 创建nginx文件挂载文件夹

  ```
  mkdir /usr/local/dockernginxdist
  ```

- 授权文件夹

  ```
  chmod 777 /usr/local/dockernginxdist
  ```

- 构建Docker镜像

  ```
  docker build -t vue_app . 
  ```

- 运行容器

  ```
  docker run -d -p 8088:80 vue_app
  ```

- 访问

  ```
  http://your_serve_ip:port
  ```
记得修改 ``` nginx.conf ```中的port
## 开发规范

- 使用Vite优化构建
- Docker镜像本地部署
- 测试和发布环境部署


项目使用Vue+Tailwind+Element搭建,采用Docker一键部署上线。
