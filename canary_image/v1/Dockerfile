FROM registry.cn-hangzhou.aliyuncs.com/launcher/alpine:latest

MAINTAINER xishengcai <cc710917049@163.com>
COPY ./bin /usr/local/bin
EXPOSE 80/tcp
RUN chmod +x /usr/local/bin/canary
ENTRYPOINT ["canary"]

