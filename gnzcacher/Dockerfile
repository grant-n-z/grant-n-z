FROM golang:1.13

RUN go get github.com/tomoyane/grant-n-z/gnzcacher

WORKDIR /go/src/github.com/tomoyane/grant-n-z/gnzcacher

ENV GOPATH $GOPATH:/go/src

RUN go build

ENTRYPOINT ["./gnzcacher"]
