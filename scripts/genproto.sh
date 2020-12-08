#!/bin/sh

set -e

proto_files=$(find ./proto -regex ".*\.\(proto\)")

for file in $proto_files; do
  echo "building proto file $file"
  protoc -I=. --go_out=. --go-grpc_out=. "$file"
done

cp -r github.com/fdymylja/cosmos-os/* ./
rm -rf github.com