#!/bin/bash

set -e

gf build main.go -a amd64 -s linux -p ./temp
gf docker main.go -p -t swcoffee/auditlimit:latest
now=$(date +"%Y%m%d%H%M")
# 以当前时间为版本号
docker tag swcoffee/auditlimit:latest swcoffee/auditlimit:$now
# docker push xyhelper/auditlimit:$now
echo "release success" $now
