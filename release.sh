#!/bin/bash

set -e

# 构建多架构二进制
gf build main.go -a amd64 -s linux -p ./temp/amd64
gf build main.go -a arm64 -s linux -p ./temp/arm64

# 使用 buildx 构建多架构镜像
docker buildx create --use
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t swcoffee/auditlimit:latest \
  --push \
  .

now=$(date +"%Y%m%d%H%M")
# 为每个架构打标签
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t swcoffee/auditlimit:$now \
  --push \
  .

echo "release success" $now
