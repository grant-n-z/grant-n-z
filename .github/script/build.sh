#!/bin/bash

set -e -u -x

cd gnzcacher
docker login -u "${DOCKER_USER}" -p "${DOCKER_PASSWORD}"
docker build -t tomohito/gnzcacher:latest .
docker push tomohito/gnzcacher:latest
