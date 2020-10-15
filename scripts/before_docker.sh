#! /bin/sh

this_dir=$(dirname "$0")
export PREFIX=/usr/local
mkdir -p $PREFIX/go/src/emaild $PREFIX/go/src/_/builds
cp -r $this_dir/../* $PREFIX/go/src/emaild
ln -s $PREFIX/go/src/emaild $PREFIX/go/src/_/builds/emaild
