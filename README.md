
WarO_Go
=========

* a Go submission for War-O as a code exercise
* this is work in progress
* code is (mostly?) idiomatic but project/GOPATH is not (yet)
* new strategy is an "API/remote" which has an implementation [here](https://github.com/codetojoy/WarO_Strategy_API_Go)

Usage:
---------

* run tests: `./test.sh`
* run tests with code coverage: `./cover.sh`
* run app: `./run.sh`

Rules:
---------

Rules are [here](Rules.md).

TODO:
---------

* The `run.sh` and `test.sh` scripts are not idiomatic.
* I tried to set up `GOPATH` etc but I just can't get it to work.
    - e.g. to use a 3rd-party package such as: https://github.com/stretchr/testify
    - Since this is functional, I have to move on, as I'm more interested in
      the language than the insufferable module/package system.
* API/remote strategy is work-in-progress 
    - remove hard-coded bits and put into JSON config 
* unit tests for other strategies
* check code coverage, esp. near console strategy 
* revisit passing pointers, esp. for Player
* more `go doc`
