#!/bin/bash

set -eu

docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -f gnzcacher/Dockerfile -t grantnz/gnzcacher:"${TAG_NAME}" .
docker push grantnz/gnzcacher:"${TAG_NAME}"

docker build -f gnzserver/Dockerfile -t grantnz/gnzserver:"${TAG_NAME}" .
docker push grantnz/gnzserver:"${TAG_NAME}"
