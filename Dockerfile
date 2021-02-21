FROM golang:1.12.0-alpine3.9

ENV GOPATH=/app

RUN mkdir -p /app/src/ad_insertion

ADD . /app/src/ad_insertion

WORKDIR /app/src/ad_insertion

RUN go build -o main .

CMD ["./main"]