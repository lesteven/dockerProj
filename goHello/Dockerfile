FROM golang:alpine

LABEL version="1.0"
LABEL decsription="go env"

COPY . /proj

WORKDIR /proj

RUN go build ./env.go

ENV VAL=world

ENTRYPOINT ["./env"]

