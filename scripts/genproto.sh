#!/bin/sh

set -e

proto_files=$(find ./module -regex ".*\.\(proto\)")

for file in $proto_files; do
  echo "building proto file $file"
  protoc -I=. -I=./third_party/proto -I=./module/abci --go_out=. --go-grpc_out=. "$file"
done

cp -r github.com/fdymylja/tmos/* ./
rm -rf github.com