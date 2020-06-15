#! /bin/sh

WORK_DIR=$(pwd); cd $(dirname $0)
PROG_DIR=$(pwd); cd $WORK_DIR

PROTOC=protoc

BASE_DIR=$PROG_DIR/..
DATA_DIR=$PROG_DIR/../../client

cd $BASE_DIR

for proto_dir in api data; do
    PROTO_DIR=proto/$proto_dir
    proto_files=$(find $PROTO_DIR -type f -name *.proto)
    for proto_file in ${proto_files[@]}; do
        $PROTOC -I=. $proto_file \
            --plugin=$PROG_DIR/protoc-gen-grpc-web \
            --js_out=import_style=commonjs:$DATA_DIR \
            --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$DATA_DIR
    done
done
