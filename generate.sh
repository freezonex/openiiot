#!/bin/bash

export RUNTIME_IDC_NAME=local


usage()
{
  echo "usage: sh generate.sh db
        db: hongzhi or openiiot"
}

case $1 in
  "hongzhi")
    go run cmd/generate/generate.go hongzhi
    ;;
  "openiiot")
      go run cmd/generate/generate.go openiiot
      ;;
  *)
    usage
    ;;
#go run cmd/generate/generate.go
esac



