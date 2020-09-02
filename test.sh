#!/bin/bash

set -e

export GOPATH=$PWD

HOME_DIR=$PWD
ROOT_DIR=$PWD/src/github.com/codetojoy

cd $ROOT_DIR/casino
go test -test.v
CASINO_RESULT=$?

cd $ROOT_DIR/player
go test -test.v
PLAYER_RESULT=$?

cd $ROOT_DIR/strategy
go test -test.v
STRATEGY_RESULT=$?

echo -e "\nTRACER casino ${CASINO_RESULT}\n"
echo -e "TRACER player ${PLAYER_RESULT}\n"
echo -e "TRACER strategy ${STRATEGY_RESULT}\n"

cd $HOME_DIR
echo "Ready."
