#!/bin/sh

set -e

build() {
    echo finding protobuf files in "$1"
    proto_files=$(find "$1" -name "*.proto")
    for file in $proto_files; do
      echo "building proto file $file"
      protoc -I=. -I=./third_party/proto --plugin /usr/bin/protoc-gen-starport --starport_out=. --go_out=. "$file"
    done
}

for dir in "$@"
do
  build "$dir"
done

cp -r github.com/fdymylja/tmos/* ./
rm -rf github.com