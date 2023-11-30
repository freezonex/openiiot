#!/bin/bash

export RUNTIME_IDC_NAME=local

usage() {
  echo "usage: sh generate.sh db
        db: openiiot"
}

case "$1" in
  "openiiot")
    go run cmd/generate/generate.go openiiot
    ;;
  *)
    usage
    exit 1
    ;;
esac
