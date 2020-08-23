#!/bin/bash

protos="
geo
hello
lbs
passenger
user
"

for proto in $protos
do
echo "Compile protobuf: " $proto
protoc --proto_path=. --micro_out=./proto/${proto}/ --go_out=./proto/${proto}/ proto/${proto}/${proto}.proto
done

echo "All complete."