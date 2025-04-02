# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.24.2-alpine AS builder

# 设置 Go 代理，加快依赖下载速度
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 设置工作目录
WORKDIR /app

# 复制 Go 依赖文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目代码
COPY . .

# 编译 Go 代码，生成可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o /fuck-the-world

# 使用轻量级的基础镜像，减小最终镜像体积
FROM alpine:latest

# 设置环境变量
ENV APP_ENV=production
# 设置工作目录
WORKDIR /

# 复制编译好的 Go 可执行文件到最终镜像
COPY --from=builder /fuck-the-world /fuck-the-world
COPY config.production.yaml /
# 暴露 8888 端口
EXPOSE 8888

# 运行应用
CMD ["/fuck-the-world","server"]
