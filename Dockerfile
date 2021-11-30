FROM marlonfan/golang:latest AS builder
WORKDIR /app
COPY . .
RUN hugo --minify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog

FROM alpine:latest
LABEL maintainer="i@marlon.life"
COPY --from=builder /app/blog /app/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' > /etc/timezone
ENTRYPOINT ["/app/blog"]
