#!/bin/bash

set -e -u -x

cd gnzcacher
docker login -u ${{ secrets.DOCKER_USER }} -p ${{ secrets.DOCKER_PASSWORD }}
docker build -t tomohito/gnzcacher:latest .
docker push tomohito/gnzcacher:latest
