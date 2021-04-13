FROM alpine:latest
LABEL maintainer="i@marlon.life"
COPY blog /usr/local/bin
ENTRYPOINT ["blog"]
