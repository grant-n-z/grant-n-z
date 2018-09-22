#!/bin/bash

set -e -u -x

go get github.com/tomoyane/grant-n-z
go get -u github.com/golang/dep/cmd/dep
go get github.com/stretchr/testify/assert

cd /go/src/github.com/tomoyane/grant-n-z

ls -la
touch out/version.txt
cat app.yaml | grep version |sed -e 's/[^0-9.]//g'

dep ensure

go build
