# Summary
Allocates memory in 1Mb blocks, up to the total requested by user.  Intended use is to validate container image memory constraints.

# Usage

allocate 4Mb of data
./golang-memtest 4

allocate 8Mb of data
./golang-memtest 8

use environment variable to specity number of Mb. 
useful when running in container
nmb=3 ./golang-memtest


# Prerequisites
* make utility (sudo apt-get install make)
* GoLang (https://fabianlee.org/2018/05/09/golang-installing-the-go-programming-language-on-ubuntu-16-04/)


# Makefile targets
* docker-run (builds golang binary, docker image, runs)
* docker-run-ok (puts 8Mb limit on docker container, allocates 4, should run OK)
* docker-run-bigmem (puts 8Mb limit on docker container, allocates 12, should FAIL)
