FROM golang:1.11

RUN go get github.com/labstack/echo
RUN mkdir -p /go/src/xcoin_rate_tracker
WORKDIR /go/src/xcoin_rate_tracker

ADD . /go/src/xcoin_rate_tracker

RUN go version
