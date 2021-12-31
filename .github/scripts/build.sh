#!/bin/bash

set -e -u -x

# gnzcacher build
now=$(date "+%Y%m%d%H%M%S")
docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -f gnzcacher/Dockerfile -t grantnz/gnzcacher:"${now}" .
docker push grantnz/gnzcacher:"${now}"
docker build -f gnzcacher/Dockerfile -t grantnz/gnzcacher:latest .
docker push grantnz/gnzcacher:latest

# gnzserver build
docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -f gnzserver/Dockerfile -t grantnz/gnzserver:"${now}" .
docker push grantnz/gnzserver:"${now}"
docker build -f gnzserver/Dockerfile -t grantnz/gnzserver:latest .
docker push grantnz/gnzserver:latest
