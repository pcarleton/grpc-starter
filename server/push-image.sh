#!/usr/bin/env bash

PROJECT_ID="cashcoach-160218"
TAG=""
CONTAINER_LABEL="cc_api_server"
IMAGE="gcr.io/$PROJECT_ID/cc-api-server-image$TAG"
INSTANCE_NAME="cashcoach-api"


docker tag "$CONTAINER_LABEL" "$IMAGE"
docker push "$IMAGE"

# Restart the vm
gcloud beta compute instances update-container $INSTANCE_NAME \
      --container-image $IMAGE
