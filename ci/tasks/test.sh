#!/bin/bash

set -e -x

service rabbitmq-server start

export GOPATH=$PWD
org=$GOPATH/src/github.com/RackHD
mkdir -p $org
cp -r utils $org/utils

pushd $org/utils
  make deps
  make test
popd
