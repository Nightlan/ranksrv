#!/bin/bash

if [[ -d "./proto/pb" ]]; then
    rm -rf ./proto/pb
fi
mkdir ./proto/pb

protoc -I ./proto/ ./proto/*.proto --go_out=plugins=grpc:./proto/pb