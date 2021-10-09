#!/bin/sh

echo "3辺の長さがa,b,cの三角形の面積を求めます"
expr='sqrt((a+b+c)*(-a+b+c)*(a-b+c)*(a+b-c))/4'
go run main.go "$expr"
