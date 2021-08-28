#!/bin/sh
set -eu
cd $(dirname $0)

echo "benchmark for Loop echo"
go test -bench=Loop main_test.go | grep /op

echo "benchmark for Join echo"
go test -bench=Join main_test.go | grep /op
