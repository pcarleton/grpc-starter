#!/usr/bin/env bash

# TODO: Bazel?
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./bin/run_server/
docker build -t cc_api_server -f Dockerfile .
