#!/bin/bash 

set -e

export GOPATH=$PWD

HOME_DIR=$PWD
ROOT_DIR=$PWD/src/github.com/codetojoy

function gen_coverage {
    go test -v -coverprofile cover.out .
    go tool cover -html=cover.out -o cover.html
}

cd $ROOT_DIR/casino
gen_coverage

cd $ROOT_DIR/player
gen_coverage

cd $HOME_DIR
find . -name "cover.html"
echo "Ready."
