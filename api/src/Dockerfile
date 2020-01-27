FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk add git && go build . && go install -v .

CMD go run github.com/ashishghogre/goapi