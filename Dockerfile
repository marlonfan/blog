ARG GO_VERSION
ARG BASE_VERSION
FROM golang:1.16.3 AS go-build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o blog

FROM alpine:latest
LABEL maintainer="i@marlon.life"
COPY --from=go-build /app/blog /app/
ENTRYPOINT ["/app/blog"]
