FROM golang:1.6-alpine
MAINTAINER Michal Rostecki <michal.rostecki@gmail.com>

RUN mkdir -p /opt/go/src/github.com/nhlfr
RUN addgroup -S go \
    && adduser -D -S -G go go

ADD . /opt/go/src/github.com/nhlfr/docker-signals-example
RUN chown -R go:go /opt/go

USER go

ENV GOPATH /opt/go

RUN go install github.com/nhlfr/docker-signals-example

CMD ["/opt/go/bin/docker-signals-example"]
