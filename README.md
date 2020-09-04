
WarO_Go
=========

* a Go submission for War-O as a code exercise
* this is work in progress
* code is (mostly?) idiomatic but project/GOPATH is not (yet)

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
* unit tests for other strategies
* check code coverage, esp. near console strategy 
* revisit passing pointers, esp. for Player
y * idiomatic `go doc`
x * idiomatic code format (use `go fmt`)
x * clean up various functions, dead code
x * revisit public functions
x * write Console strategy
x * remove Bid and add Offer to Player 
x * use JSON for configuration
x * use interface for strategies
x * consider new `players` package for functions on a plurality of players
