#!/usr/bin/env bash

curl -L -O https://github.com/gohugoio/hugo/releases/download/v0.64.1/hugo_extended_0.64.1_Linux-64bit.tar.gz
tar -xzf hugo_extended_0.64.1_Linux-64bit.tar.gz

curl -L -O ftp://ftp.pbone.net/mirror/li.nux.ro/download/nux/dextop/el6/x86_64/chrome-deps-stable-3.11-1.x86_64.rpm
rpm -i --badreloc --noscripts --relocate /opt/google/chrome=$HOME chrome-deps-stable-3.11-1.x86_64.rpm
export LD_LIBRARY_PATH=$HOME/lib:$LD_LIBRARY_PATH

./hugo -v