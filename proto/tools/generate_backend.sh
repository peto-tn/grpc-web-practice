#!/bin/sh

WORK_DIR=$(pwd); cd $(dirname $0)
PROG_DIR=$(pwd); cd $WORK_DIR

PROTOC=protoc

BASE_DIR=$PROG_DIR/..
DATA_DIR=$PROG_DIR/../../backend/app

cd $BASE_DIR

for proto_dir in api data; do
    PROTO_DIR=proto/$proto_dir
    proto_files=$(find $PROTO_DIR -type f -name *.proto)
    mkdir -p $DATA_DIR/$PROTO_DIR
    for proto_file in ${proto_files[@]}; do
        service=$(basename ${proto_file%.*})
        $PROTOC --go_out=plugins=grpc:$DATA_DIR $PROTO_DIR/$service.proto
        sed -i "" -e "s/proto\/data/blog\/app\/proto\/data/" $DATA_DIR/$PROTO_DIR/$service.pb.go
    done
done
