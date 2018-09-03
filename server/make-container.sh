#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./bin/run_server/
docker build -t $IMAGE_NAME -f Dockerfile .
