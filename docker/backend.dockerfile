FROM golang:1.13.7-alpine3.11

WORKDIR /app

ENV GOPATH "${GOPATH}:/app"

COPY src/iio /app/src/iio

RUN go build iio

CMD ["./iio"]
