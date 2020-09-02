#!/bin/bash 

set -e

export GOPATH=$PWD

HOME_DIR=$PWD
ROOT_DIR=$PWD/src/github.com/codetojoy

function gen_coverage {
    WORKING_DIR=$1
    cd $ROOT_DIR/$WORKING_DIR
    go test -v -coverprofile cover.out .
    go tool cover -html=cover.out -o cover.html
    cd $ROOT_DIR
}

gen_coverage "casino"
gen_coverage "player"
gen_coverage "strategy"


cd $HOME_DIR
find . -name "cover.html"
echo "Ready."
