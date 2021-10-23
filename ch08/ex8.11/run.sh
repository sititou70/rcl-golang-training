#!/bin/sh
set -eux

go run main.go \
  https://golang.org/ \
  https://www.google.com/ \
  https://github.com/ \
  https://www.youtube.com/ \
  https://twitter.com/home \
  https://www.facebook.com/
