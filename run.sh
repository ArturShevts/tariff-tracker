#!/bin/bash

case "$1" in
  "backend")
    cd apps/backend && go run cmd/server/main.go
    ;;
  "build-backend")
    cd apps/backend && go build -o ../../bin/backend cmd/server/main.go
    ;;
  *)
    echo "Usage: ./run.sh [backend|build-backend]"
    exit 1
    ;;
esac
