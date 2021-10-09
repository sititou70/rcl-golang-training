#!/bin/sh
set -eu

curl https://www.w3.org/TR/2006/REC-xml11-20060806 | go run main.go "a[href='/']"
