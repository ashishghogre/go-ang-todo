FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN apk add git && go get -d -v . && go install -v .

CMD app