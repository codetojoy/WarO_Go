#!/bin/bash

CONFIG_FILE=$1
if [ ! -f "$CONFIG_FILE" ]; then
    echo "could not find config file: '$CONFIG_FILE'"
    exit -1
fi

export GOPATH=$PWD

go run src/github.com/codetojoy/runner.go $CONFIG_FILE 
