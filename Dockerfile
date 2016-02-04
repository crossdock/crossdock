FROM golang
ENV GO15VENDOREXPERIMENT 1
ADD . /go/src/github.com/yarpc/crossdock
RUN go install github.com/yarpc/crossdock
ENTRYPOINT /go/bin/crossdock
