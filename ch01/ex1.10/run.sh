#!/bin/sh
set -eu
cd $(dirname $0)

go run main.go https://golang.org http://gopl.io https://godoc.org
