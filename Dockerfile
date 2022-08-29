FROM marlonfan/golang:latest AS builder
WORKDIR /app
COPY . .
RUN hugo --minify
COPY root.txt /app/public
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog

FROM alpine:latest
LABEL maintainer="i@marlon.life"
COPY --from=builder /app/blog /app/
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
ENTRYPOINT ["/app/blog"]
