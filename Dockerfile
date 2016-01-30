FROM golang
ENV GO15VENDOREXPERIMENT 1
RUN curl --location --silent --show-error --fail \
        https://github.com/Barzahlen/waitforservices/releases/download/v0.3/waitforservices \
        > /usr/local/bin/waitforservices && \
    chmod +x /usr/local/bin/waitforservices
ADD . /go/src/github.com/yarpc/xlang
RUN go install github.com/yarpc/xlang
CMD ["/go/bin/xlang"]
