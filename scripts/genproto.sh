#!/bin/sh

set -e

proto_files=$(find ./core -regex ".*\.\(proto\)")

for file in $proto_files; do
  echo "building proto file $file"
  protoc -I=. -I=./third_party/proto -I=./core/abci --go_out=. --go-grpc_out=. "$file"
done

proto_files=$(find ./x -regex ".*\.\(proto\)")
for file in $proto_files; do
  echo "building proto file $file"
  protoc -I=. -I=./third_party/proto -I=./core/abci --go_out=. --go-grpc_out=. "$file"
done

cp -r github.com/fdymylja/tmos/* ./
rm -rf github.com