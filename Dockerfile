FROM --platform=$BUILDPLATFORM alpine:latest

WORKDIR /cloudreve
COPY cloudreve ./cloudreve

RUN apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime \
    && echo "Asia/Taipei" > /etc/timezone \
    && chmod +x ./cloudreve \
    && mkdir -p /data/aria2 \
    && chmod -R 766 /data/aria2

EXPOSE 5212
VOLUME ["/cloudreve/uploads", "/cloudreve/avatar", "/data"]

ENTRYPOINT ["./cloudreve"]