#!/bin/bash

DIR=$( cd "$(dirname "${BASH_SOURCE[0]}")/.." ; pwd -P )

THEAPP_CONFIG_PATH=$DIR/test/input.json \
THEAPP_OUTPUT_PATH=$DIR/test/output.png \
go run $DIR/main.go 