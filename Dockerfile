FROM golang
ADD . /go/src/github.com/yarpc/xlang
RUN go install github.com/yarpc/xlang
ENTRYPOINT /go/bin/xlang
