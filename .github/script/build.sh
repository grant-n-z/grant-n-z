#!/bin/bash

set -e -u -x

# gnzcacher build
cd gnzcacher
cat grant_n_z_cacher.yaml| grep version | sed 's/^[ \t]*//' | sed 's/version://' | sed 's/^[ \t]*//' > version
ver=`cat version`
ver=`cat version`

GOOS=linux GOARCH=amd64 go build

docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -t grantnz/gnzcacher:"${ver}" -t grantnz/gnzcacher:latest .
docker push grantnz/gnzcacher

# gnzserver build
cd ../gnzserver
cat grant_n_z_server.yaml| grep version | sed 's/^[ \t]*//' | sed 's/version://' | sed 's/^[ \t]*//' > version
ver=`cat version`

GOOS=linux GOARCH=amd64 go build

docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -t grantnz/gnzserver:"${ver}" -t grantnz/gnzserver:latest .
docker push grantnz/gnzserver
