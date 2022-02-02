#!/bin/sh
#shellcheck disable=SC1091
. .env

if [ "$#" = 1 ]; then
  docker compose run --rm "golang-${ENVIRONMENT}" "air" "-c" ".air.toml"
else
  if [ "$ENVIRONMENT" = 'prod' ]; then
    docker compose build --force-rm --no-cache
    docker compose up -d
  else
    docker compose run --rm "golang-${ENVIRONMENT}" "go" "run" "main.go"
  fi
fi
