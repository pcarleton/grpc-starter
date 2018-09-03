#!/usr/bin/env bash
TAG=""
IMAGE="gcr.io/$PROJECT_ID/$IMAGE_NAME$TAG"

docker tag "$IMAGE_NAME" "$IMAGE"
docker push "$IMAGE"
