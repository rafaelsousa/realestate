#!/bin/bash

# Generates the proto components

protoc --proto_path=proto --proto_path=third_party/proto --js_out=import_style=commonjs,binary:ui/gen \
    property/genesis.proto            \
    property/owner.proto              \
    property/property.proto           \
    property/query.proto              \
    property/tx.proto                 \
    realestate/genesis.proto          \
    realestate/query.proto            \
    realestate/tx.proto
