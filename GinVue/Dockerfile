FROM golang:1.21.8-alpine

# 维护人员信息
MAINTAINER JaeHua

# 工作目录
WORKDIR $GOPATH/src/gin

# 将本地内容添加到镜像指定目录
COPY . $GOPATH/src/gin

# 设置开启go mod
RUN go env -w GO111MODULE=auto

# 设置go代理
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 构建go应用
RUN go build

# 暴露端口
EXPOSE 3344

# 镜像默认入口命令
ENTRYPOINT ["./GinVue"]