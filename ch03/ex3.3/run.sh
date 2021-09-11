#!/bin/sh
set -eu
cd $(dirname $0)

go run main.go eggCase > eggCase.svg
go run main.go mogul > mogul.svg
go run main.go saddle > saddle.svg
