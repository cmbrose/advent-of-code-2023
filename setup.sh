#! /bin/bash

dest="$1"
src="${2:-template}"

cp -r $src $dest

cd $dest

code *