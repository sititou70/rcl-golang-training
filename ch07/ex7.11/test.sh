#!/bin/sh
set -eu
cd $(dirname $0)

curl "localhost:8000/create?item=shirt&price=100"
curl "localhost:8000/update?item=shirt&price=200"
curl "localhost:8000/delete?item=shoes"

expected="socks: \$5
shirt: \$200"
if [ "$(curl 'localhost:8000/list')" != "$expected" ]; then
  echo "fail"
  exit 1
fi

echo "ok"
