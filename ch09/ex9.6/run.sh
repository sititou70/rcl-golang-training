#!/bin/sh
set -eux
cd $(dirname $0)

export TIME="%e seconds"

MAX_GOMAXPROCS=20

for gomaxprocs in $(seq $MAX_GOMAXPROCS); do
  GOMAXPROCS=$gomaxprocs time go run main.go $MAX_GOMAXPROCS >/dev/null
done
