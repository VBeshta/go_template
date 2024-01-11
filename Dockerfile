FROM golang:1.21

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go clean --mdcache
RUN go build -o main .
CMD ["/app/main"]