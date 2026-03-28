#!/usr/bin/env bash

set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo "Usage: $0 <dockerhub-username> [image-name]" >&2
  echo "Example: $0 johndoe go-api-app" >&2
  exit 1
fi

DOCKERHUB_USERNAME="$1"
IMAGE_NAME="${2:-go-api-app}"
FULL_IMAGE_NAME="${DOCKERHUB_USERNAME}/${IMAGE_NAME}"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

if [[ ! -f "${PROJECT_ROOT}/Dockerfile" ]]; then
  echo "ERROR: Dockerfile not found at ${PROJECT_ROOT}/Dockerfile"
  exit 1
fi

SHA_TAG="$(git -C "${PROJECT_ROOT}" rev-parse --short HEAD)"

if ! command -v docker >/dev/null 2>&1; then
  echo "ERROR: docker is not installed or not in PATH"
  exit 1
fi

if ! docker info >/dev/null 2>&1; then
  echo "ERROR: Docker daemon is not running"
  exit 1
fi

if ! docker buildx version >/dev/null 2>&1; then
  echo "ERROR: docker buildx is not available"
  exit 1
fi

if ! docker buildx inspect render-builder >/dev/null 2>&1; then
  docker buildx create --name render-builder --driver docker-container --use >/dev/null
else
  docker buildx use render-builder >/dev/null
fi

echo "Building and pushing linux/amd64 image for Render..."
echo "Image: ${FULL_IMAGE_NAME}"
echo "Tags: latest, ${SHA_TAG}"

docker buildx build \
  --platform linux/amd64 \
  --file "${PROJECT_ROOT}/Dockerfile" \
  --tag "${FULL_IMAGE_NAME}:latest" \
  --tag "${FULL_IMAGE_NAME}:${SHA_TAG}" \
  --push \
  "${PROJECT_ROOT}"

echo "Done."
echo "Pushed: ${FULL_IMAGE_NAME}:latest"
echo "Pushed: ${FULL_IMAGE_NAME}:${SHA_TAG}"