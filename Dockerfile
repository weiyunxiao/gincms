FROM golang:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache tzdata

ENV TZ Asia/Shanghai

WORKDIR /www
# 复制编译阶段编译出来的运行文件到目标目录
COPY gincms .
# 复制编译阶段里的config文件夹到目标目录
COPY config.yaml .
# 将时区设置为东八区
RUN chmod +x gincms
# 需暴露的端口
EXPOSE 8066
# 可外挂的目录
VOLUME ["/www/log"]
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./gincms"]