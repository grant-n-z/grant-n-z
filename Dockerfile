FROM golang:1.9.4 

ENV GOPATH $GOPATH:/go/src 

RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

RUN mkdir /go/src/revel-api

EXPOSE 9000
