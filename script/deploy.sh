#!/usr/bin/env bash
#
# deploy.sh - 部署生产环境。
set -euo pipefail

cd "$(dirname "$0")/.."

IMAGE_REPOSITORY="${IMAGE_REPOSITORY:-ghcr.io/hcd233/aris-api-tmpl}"
BRANCH_NAME="${BRANCH_NAME:-$(git rev-parse --abbrev-ref HEAD)}"
IMAGE_TAG="${IMAGE_TAG:-$(echo "${BRANCH_NAME}" | tr '/' '-')}"
COMPOSE_FILE="${COMPOSE_FILE:-docker/docker-compose-single.yml}"
SERVICE_NAME="${SERVICE_NAME:-aris-api-tmpl}"

printf '\033[1;36mDeploying production environment (branch: %s, image: %s:%s)...\033[0m\n' "${BRANCH_NAME}" "${IMAGE_REPOSITORY}" "${IMAGE_TAG}"

if [ "${SKIP_GIT_PULL:-false}" != "true" ]; then
    printf '\033[1;36mPulling the latest code...\033[0m\n'
    git fetch --prune origin
    git pull --ff-only origin "${BRANCH_NAME}"
fi

printf '\033[1;32mPulling Docker image...\033[0m\n'
docker pull "${IMAGE_REPOSITORY}:${IMAGE_TAG}"

printf '\033[1;34mStarting services with docker compose...\033[0m\n'
export IMAGE_REPOSITORY IMAGE_TAG
docker compose -f "${COMPOSE_FILE}" up -d

if [ "${PRUNE_IMAGES:-false}" = "true" ]; then
    printf '\033[1;31mPruning unused Docker images...\033[0m\n'
    docker image prune -f
fi

printf '\033[1;33mRecent logs for %s...\033[0m\n' "${SERVICE_NAME}"
if [ "${FOLLOW_LOGS:-false}" = "true" ]; then
    docker logs -f "${SERVICE_NAME}" --details --tail 25
else
    docker logs "${SERVICE_NAME}" --details --tail 25
fi
