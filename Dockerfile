FROM golang
ENV GO15VENDOREXPERIMENT 1
ADD . /go/src/github.com/yarpc/xlang
RUN go install github.com/yarpc/xlang
ENTRYPOINT /go/bin/xlang
