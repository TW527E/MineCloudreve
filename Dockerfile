FROM alpine:latest

WORKDIR /cloudreve
COPY cloudreve ./cloudreve

RUN apk update
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime
RUN echo "Asia/Taipei" > /etc/timezone
RUN chmod +x ./cloudreve
RUN mkdir -p /data/aria2
RUN chmod -R 766 /data/aria2

EXPOSE 5212
VOLUME ["/cloudreve/uploads", "/cloudreve/avatar", "/data"]

ENTRYPOINT ["./cloudreve"]