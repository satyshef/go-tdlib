#!/bin/bash

apt install -y cmake build-essential libssl-dev zlib1g-dev gperf ccache libreadline-dev git

git clone git@github.com:tdlib/td.git --depth 1
cd td
mkdir build
cd build
cmake -DCMAKE_BUILD_TYPE=Release ..
cmake --build . -- -j1
make install
