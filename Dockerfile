# 使用 golang 的基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod tidy

# 复制所有文件到工作目录
COPY . .

# 编译可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# 使用 scratch 作为基础镜像
FROM scratch

# 将可执行文件复制到镜像中
COPY --from=0 /app/app /

# 暴露端口
EXPOSE 9551
EXPOSE 9552
EXPOSE 9553
EXPOSE 9554
EXPOSE 9651
EXPOSE 9652
EXPOSE 9653
EXPOSE 9654

# 运行可执行文件
ENTRYPOINT ["/app"]