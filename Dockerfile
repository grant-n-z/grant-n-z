FROM golang:1.9.4

ENV GOPATH $GOPATH:/go/src
ENV DB_SOURCE="root:root@tcp(docker.for.mac.localhost:3306)/auth_server?charset=utf8&parseTime=True"

RUN go get github.com/tomoyane/grant-n-z
RUN cd /go/src/github.com/tomoyane/grant-n-z

CMD dep ensure
CMD revel run grant-n-z
