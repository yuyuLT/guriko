#!/bin/sh

if [ -e go.mod ]; then
  go mod tidy
else
  go mod init "$MODULE_PATH"
fi

exec "$@"
