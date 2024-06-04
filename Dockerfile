# Docker builder for Golang
FROM golang as builder

WORKDIR /go/src/github.com/thcdrt/grands-buffets-notifier
COPY . .
RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


# Docker run Golang app
FROM scratch

WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/thcdrt/grands-buffets-notifier .
CMD ["./app"]