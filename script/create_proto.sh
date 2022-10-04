#!/bin/bash

SERVICE_NAME=$1
SERVICE_NAME_UPPER="$(tr '[:lower:]' '[:upper:]' <<< ${SERVICE_NAME:0:1})${SERVICE_NAME:1}"
PROTO_PATH="$SERVICE_NAME/pkg/${SERVICE_NAME}proto"

echo "Creating $PROTO_PATH"

mkdir -pv $PROTO_PATH
cat <<EOF > $PROTO_PATH/$SERVICE_NAME.proto
syntax = "proto3";

option go_package = "../${SERVICE_NAME}proto";

package $SERVICE_NAME;

service $SERVICE_NAME_UPPER {
  
}
EOF

echo "Created $PROTO_PATH"