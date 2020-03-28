#!/bin/bash

set -e -u -x

go test -coverprofile=coverage.out -v ./...
