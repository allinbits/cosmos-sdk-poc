#!/bin/sh

set -e

proto_files=$(find ./apis -regex ".*\.\(proto\)")

for file in $proto_files; do
  echo "building proto file $file"
  protoc -I=. -I=./third_party/proto -I=./apis/abci --go_out=. --go-grpc_out=. "$file"
done

cp -r github.com/fdymylja/tmos/* ./
rm -rf github.com