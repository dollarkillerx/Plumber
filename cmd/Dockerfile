#FROM alpine:latest
#WORKDIR /app
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
#    && apk add mariadb-client
#
#COPY ./cmd/plumber /app
#COPY ./configs/config.yaml /app/configs/config.yaml
#CMD ["/app/plumber"]

FROM ubuntu:20.04
WORKDIR /app
RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list \
    && apt-get update \
    && apt-get install mysql-client -y

COPY ./cmd/plumber /app
COPY ./configs/config.yaml /app/configs/config.yaml
CMD ["/app/plumber"]