#!/bin/bash

set -eu

docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -t grantnz/gnzcacher:"${TAG_NAME}" .
docker push grantnz/gnzcacher:"${TAG_NAME}"

docker build -t grantnz/gnzserver:"${TAG_NAME}" .
docker push grantnz/gnzserver:"${TAG_NAME}"
