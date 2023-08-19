FROM golang:1.19.3-alpine3.16

COPY . /app/src
WORKDIR /app/src

RUN go build .

CMD ["/app/src/go_simple_chat_app_api"]
