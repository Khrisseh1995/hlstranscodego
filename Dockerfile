FROM golang:1.12.0-alpine3.9

ENV GOPATH=/app

RUN mkdir -p /app/src/rest_api

ADD . /app/src/rest_api

WORKDIR /app/src/rest_api

RUN go build -o main .

# RUN ls

# CMD ["./main"]