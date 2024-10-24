FROM --platform=$BUILDPLATFORM alpine:latest

WORKDIR /cloudreve
COPY cloudreve ./cloudreve

RUN apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Taipei /etc/localtime \
    && echo "Asia/Taipei" > /etc/timezone \
    && apk add --no-cache ffmpeg \
    && apk add --no-cache vips vips-tools \
    && apk add --no-cache libreoffice \
    && apk add --no-cache libraw libraw-tools \
    && chmod +x ./cloudreve \
    && mkdir -p /data/aria2 \
    && chmod -R 766 /data/aria2

EXPOSE 5212
VOLUME ["/cloudreve/uploads", "/cloudreve/avatar", "/data"]

ENTRYPOINT ["./cloudreve"]