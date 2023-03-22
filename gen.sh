#!/bin/bash

# 更新依赖包
go get -v github.com/buexplain/netsvr-protocol@latest

# 将依赖包拉入当前目录
go mod vendor

# 删除旧的代码
mkdir -p ./protocol
rm -rf ./protocol/*

#生成新的代码
# https://protobuf.dev/reference/go/go-generated/

# shellcheck disable=SC2046
root_dir="$(pwd)"
proto_path="$root_dir/vendor/github.com/buexplain/netsvr-protocol"
match="$proto_path/*.proto"
for file in $match; do
  proto="$(basename "$file")"
  echo "protoc --go_out=$root_dir/ --proto_path=$proto_path $proto"
  protoc --go_out="$root_dir" --proto_path="$proto_path" "$proto"
  # shellcheck disable=SC2181
  if [ $? != 0 ]; then
    # shellcheck disable=SC2162
    read
    exit $?
  fi
done

# 清空依赖信息
rm -r vendor