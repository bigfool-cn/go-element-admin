FROM golang:alpine as builder

# MAINTAINER
LABEL name="go-element-admin"
LABEL version="1.0.1"
LABEL author="bigfool <1063944784@qq.com>"
LABEL maintainer="bigfool <1063944784@qq.com>"
LABEL description="go-element-admin application"

# 开启Go Module, 设置GO Proxy代理
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

# 新建项目目录
RUN mkdir -p /go-element-admin

# 指定工作目录
WORKDIR /go-element-admin

# 复制源代码到工作目录
COPY . .

# 删除旧的可执行文件
RUN rm -rf /go-element-admin/server

# 设置操作系统, 操作系统构架
RUN GOOS=linux GOARCH=amd64

RUN go build -o server .

# 添加可执行权限
RUN chmod +x /go-element-admin/server


FROM alpine

# 复制builder相关文件到基础镜像alipine
COPY --from=builder /go-element-admin/configs/config.yaml /go-element-admin/configs/config.yaml
COPY --from=builder /go-element-admin/server /go-element-admin

# 设置时区
RUN apk add -U tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/lcoaltime \
    && echo 'Asia/Shanghai' > /etc/timezone

ENV TZ=Asia/Shanghai

# 创建日志目录
RUN mkdir -p /go-element-admin/logs

# 新建一个用户www 并设置项目目录用户组
RUN adduser -D -H www \
    && chown -R www /go-element-admin

# 执行用户
USER www

WORKDIR /go-element-admin

EXPOSE 8001

ENTRYPOINT ["./server"]
