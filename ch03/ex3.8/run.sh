#!/bin/sh
set -eu
cd $(dirname $0)

go run main.go complex64 > complex64.png
go run main.go complex128 > complex128.png
go run main.go big.Float > big.Float.png
#go run main.go big.Rat > big.Rat.png
