# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest

# 在容器中创建工作目录
WORKDIR /app

# 将项目中的所有文件拷贝到容器的工作目录
COPY . .

# 将配置文件拷贝到容器内的 /app/config 目录
COPY config/ /app/config/

# 构建 Go 应用
RUN go build -o bot .

# 设置启动命令
CMD ["./bot"]
