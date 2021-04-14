FROM marlonfan/golang:latest AS go-build
WORKDIR /app
COPY . .
RUN hugo --minify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog

FROM alpine:latest
LABEL maintainer="i@marlon.life"
COPY --from=go-build /app/blog /app/
ENTRYPOINT ["/app/blog"]
