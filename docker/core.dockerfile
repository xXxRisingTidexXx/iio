FROM golang:1.14-alpine AS base

WORKDIR /go

COPY cmd /go/src/cmd

RUN adduser -D -g "" iio && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/core


FROM alpine:3.11 AS main

WORKDIR /app

COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /go/core /app/core

USER iio

CMD ["./core"]
