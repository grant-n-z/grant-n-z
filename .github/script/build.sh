#!/bin/bash

set -e -u -x

cd gnzcacher
docker login docker.pkg.github.com -u "${DOCKER_USER}" -p "${GITHUB_TOKEN}"
docker build -t docker.pkg.github.com/tomoyane/grant-n-z/gnzcacher:latest .
docker push docker.pkg.github.com/tomoyane/grant-n-z/gnzcacher:latest
