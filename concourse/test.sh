#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

mkdir /go/src/github.com/tomoyane/

mv repository grant-n-z
cp -r grant-n-z /go/src/github.com/tomoyane/

cd /go/src/github.com/tomoyane/grant-n-z

dep ensure

go test -v github.com/tomoyane/grant-n-z/test/...
