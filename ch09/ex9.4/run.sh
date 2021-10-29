#!/bin/sh
set -eu
cd $(dirname $0)

run() {
  echo $1 goroutine
  go run main.go $1
}

run 1
run 10
run 100
run 1000
run 10000
run 100000
run 1000000
run 10000000
run 100000000
run 1000000000
