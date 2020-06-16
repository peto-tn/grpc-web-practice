#!/bin/bash

WORK_DIR=$(pwd); cd $(dirname $0)
PROG_DIR=$(pwd); cd $WORK_DIR

BASE_DIR=$PROG_DIR/../../../app
DATA_DIR=$PROG_DIR/../../../app/library/test/mock

cd $BASE_DIR

for target_dir in proto/api application domain/service domain/repository; do
    SEARCH_DIR=$target_dir
    src_files=$(find $SEARCH_DIR -type f -name *.go -not -name '*_test.go')
    for src_file in ${src_files[@]}; do
        if [ "$src_file" = "proto/api/api.pb.go" ]; then continue; fi
        echo $src_file
        mkdir -p $DATA_DIR/$(dirname $src_file)
        mockgen -source $src_file > $DATA_DIR/$src_file
    done
done

