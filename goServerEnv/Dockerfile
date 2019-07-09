FROM golang:alpine

LABEL version="1.0"
LABEL decsription="go server env"

COPY . /proj

WORKDIR /proj

RUN go build ./server.go

EXPOSE 8080

ENV VAL=world

ENTRYPOINT ["./server"]

