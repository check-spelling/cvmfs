#!/bin/sh

make clean
make
# Don't strip debug symbols
# strip -S lib/.libs/libcurl.a
make install
