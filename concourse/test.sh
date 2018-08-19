#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

cp -r repository /go/src/

cd /go/src/repository

dep ensure

go test test