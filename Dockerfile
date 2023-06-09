FROM golang:1.20-alpine

ADD . /go/src/bit-exporter
WORKDIR /go/src/bit-exporter
RUN apk add --no-cache make
RUN go get .
RUN make build-static
