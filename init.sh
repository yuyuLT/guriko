#!/bin/sh

find . -type f -name "*.sh" -exec chmod a+x "{}" +
docker compose build --force-rm --no-cache
