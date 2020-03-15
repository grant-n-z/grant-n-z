#!/bin/bash

set -e -u -x

cd gnzserver && go build && cd ..
cd gnzcacher && go build && cd ..
