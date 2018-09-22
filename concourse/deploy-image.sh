#!/bin/bash

set -e -u -x

go get github.com/tomoyane/grant-n-z
go get -u github.com/golang/dep/cmd/dep
go get github.com/stretchr/testify/assert

cd /go/src/github.com/tomoyane/grant-n-z

touch ../version/version.txt
cat app.yaml | grep version |sed -e 's/[^0-9.]//g' >> ../version/version.txt

dep ensure

go build
