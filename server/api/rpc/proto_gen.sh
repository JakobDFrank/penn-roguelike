#!/bin/bash
cd "$(dirname "$0")" || exit 1
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./*.proto