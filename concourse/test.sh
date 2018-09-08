#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

cp -r repository /go/src/

cd /go/src/repository

dep ensure

pwd
ls -la
go test -v /go/src/repository/test/...