#!/bin/bash

SERVICE_NAME=$1
SERVICE_NAME_UPPER="$(tr '[:lower:]' '[:upper:]' <<< ${SERVICE_NAME:0:1})${SERVICE_NAME:1}"
PROTO_PATH="$SERVICE_NAME/internal/libproto"

echo "Creating $PROTO_PATH"

mkdir -pv $PROTO_PATH
cat <<EOF > $PROTO_PATH/$SERVICE_NAME.proto
syntax = "proto3";

option go_package = "../libproto";

package $SERVICE_NAME;

service $SERVICE_NAME_UPPER {
  
}
EOF

echo "Created $PROTO_PATH"