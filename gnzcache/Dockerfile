FROM golang:1.9.4

RUN go get github.com/golang/dep/cmd/dep && \
    go get github.com/tomoyane/grant-n-z

WORKDIR /go/src/github.com/tomoyane/grant-n-z

ENV GOPATH $GOPATH:/go/src

RUN dep ensure && \
    go build

ENTRYPOINT ["./grant-n-z"]