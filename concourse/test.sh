#!/bin/bash

set -e -u -x

go get github.com/tomoyane/grant-n-z
go get -u github.com/golang/dep/cmd/dep

cd /go/src/github.com/tomoyane/grant-n-z

dep ensure

go test -v github.com/tomoyane/grant-n-z/test/...
