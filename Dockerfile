FROM golang:1.9.4

RUN git clone https://github.com/tomoyane/grant-n-z.git && \
    go get github.com/golang/dep/cmd/dep && \
    mkdir /go/src/github.com/tomoyane && \
    cp -rp grant-n-z /go/src/github.com/tomoyane/

WORKDIR /go/src/github.com/tomoyane/grant-n-z

ENV GOPATH $GOPATH:/go/src

RUN dep ensure && \
    go build && \
    ./grant-n-z
