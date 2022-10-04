#!/bin/bash

SERVICE_NAME=$1
PROTO_PATH="$SERVICE_NAME/pkg/${SERVICE_NAME}proto"

export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=$PROTO_PATH --go-grpc_out=$PROTO_PATH $PROTO_PATH/*.proto
