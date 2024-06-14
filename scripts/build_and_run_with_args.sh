#!/bin/sh
cd ..
pwd
make build
./build/tmp/bin/opertizen "$@"