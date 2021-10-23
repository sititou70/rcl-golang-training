#!/bin/sh
set -eux

for i in $(seq 20); do
  time go run main.go $i >/dev/null
done
