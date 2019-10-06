FROM golang:alpine

RUN apk add git

WORKDIR /go/src/app

COPY . .

RUN rm -f *.db*

RUN go get -d -v .
RUN go install -v .

CMD app