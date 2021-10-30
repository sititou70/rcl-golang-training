#!/bin/sh
set -eux
cd $(dirname $0)

go <sample.jpg \
  run main.go -t jpeg |
  go run main.go -t png |
  go run main.go -t jpeg |
  go run main.go -t gif |
  go run main.go -t png |
  go run main.go -t png |
  go run main.go -t gif |
  go run main.go -t gif |
  go run main.go -t jpeg >TEMP_out.jpg
