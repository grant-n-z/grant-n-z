#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

mkdir /go/src/github.com/tomoyane/

cp -r repository /go/src/github.com/tomoyane/

mv /go/src/github.com/tomoyane/repository /go/src/github.com/tomoyane/grant-n-z
cd /go/src/github.com/tomoyane/grant-n-z

dep ensure

echo $GOPATH

go test -v github.com/tomoyane/grant-n-z/test/...
