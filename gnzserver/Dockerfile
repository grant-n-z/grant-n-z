FROM golang:1.13

RUN go get github.com/tomoyane/grant-n-z/gnzserver

WORKDIR /go/src/github.com/tomoyane/grant-n-z/gnzserver

ENV GOPATH $GOPATH:/go/src

RUN go build

ENTRYPOINT ["./gnzserver"]
