#!/bin/bash

set -e -u -x

# gnzcacher build
cd gnzcacher
cat grant_n_z_cacher.yaml| grep version | sed 's/^[ \t]*//' | sed 's/version://' | sed 's/^[ \t]*//' > version
ver=`cat version`

GOOS=linux GOARCH=amd64 go build

docker login docker.pkg.github.com -u "${DOCKER_USER}" -p "${GITHUB_TOKEN}"
docker build -t docker.pkg.github.com/tomoyane/grant-n-z/gnzcacher:"${ver}" -t docker.pkg.github.com/tomoyane/grant-n-z/gnzcacher:latest .
docker push docker.pkg.github.com/tomoyane/grant-n-z/gnzcacher:"${ver}"
docker push docker.pkg.github.com/tomoyane/grant-n-z/gnzcacher:latest

# gnzserver build
cd ../gnzserver
cat grant_n_z_server.yaml| grep version | sed 's/^[ \t]*//' | sed 's/version://' | sed 's/^[ \t]*//' > version
ver=`cat version`

GOOS=linux GOARCH=amd64 go build

docker login docker.pkg.github.com -u "${DOCKER_USER}" -p "${GITHUB_TOKEN}"
docker build -t docker.pkg.github.com/tomoyane/grant-n-z/gnzserver:"${ver}" -t docker.pkg.github.com/tomoyane/grant-n-z/gnzserver:latest .
docker push docker.pkg.github.com/tomoyane/grant-n-z/gnzserver:"${ver}"
docker push docker.pkg.github.com/tomoyane/grant-n-z/gnzserver:latest
