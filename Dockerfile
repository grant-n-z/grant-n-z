FROM golang:1.9.4

ENV GOPATH $GOPATH:/go/src
ENV DB_HOST="db"
ENV DB_NAME="auth_server"
ENV DB_USER="root"
ENV DB_PASS="root"
ENV DB_PORT="3306"

RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql

RUN mkdir /go/src/auth-server

CMD revel run revel-performance
