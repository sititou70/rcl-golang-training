#!/bin/sh
cd $(dirname $0)
set -eux

go run clockwall/clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
