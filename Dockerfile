FROM golang:1.9.4

ENV GOPATH $GOPATH:/go/src
ENV DB_HOST="docker.for.mac.localhost"
ENV DB_NAME="auth_server"
ENV DB_USER="root"
ENV DB_PASS="root"
ENV DB_PORT="3306"

RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel && \
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/satori/go.uuid && \
    go get gopkg.in/go-playground/validator.v9 && \
    go get github.com/lestrrat/go-test-mysqld && \
    go get gopkg.in/DATA-DOG/go-sqlmock.v1
    

RUN mkdir /go/src/grant-n-z

COPY . /go/src/grant-n-z

CMD revel run grant-n-z
