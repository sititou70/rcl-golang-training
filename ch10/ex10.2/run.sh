#!/bin/sh
set -eux
cd $(dirname $0)

echo "tar test"
go <ch09.tar run main.go

echo "zip test"
go <ch09.zip run main.go
