#!/bin/bash

set -e -u -x

touch ../out/version.txt

ls -la

cat app.yaml | grep version |sed -e 's/[^0-9.]//g' > ../out/version.txt

go get -u github.com/golang/dep/cmd/dep

dep ensure
go build