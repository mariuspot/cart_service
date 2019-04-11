#! /usr/bin/env bash
BASEDIR=$(dirname "$0")
protoc -I server --go_out=plugins=grpc:$BASEDIR/../pkg/api --proto_path=$BASEDIR/../pkg/api $BASEDIR/../pkg/api/*.proto