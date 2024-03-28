FROM golang:alpine

COPY ./ /src
WORKDIR /src

RUN go build -v .
CMD [ "/src/go_simple_website" ]
