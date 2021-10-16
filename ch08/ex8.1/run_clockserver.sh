#!/bin/sh
cd $(dirname $0)
set -eux

TZ=US/Eastern go run clock2/clock2.go -port 8010 &
TZ=Asia/Tokyo go run clock2/clock2.go -port 8020 &
TZ=Europe/London go run clock2/clock2.go -port 8030 &
wait
