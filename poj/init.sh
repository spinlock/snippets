#!/bin/bash

id=$1; shift

if [ "x$id" == "x" ]; then
    echo "invalid id"
    exit 1
fi

rm -rf $id
mkdir -p $id && cd $id && wget -p -k http://poj.org/problem?id=$id

if [ $? -ne 0 ]; then
    echo $?
    exit 1
fi

cat << EOF > Makefile
all:
	cat input | go run main.go

cc:
	clang++ -O2 code.cc -o /tmp/pojcode.out && cat input | /tmp/pojcode.out
EOF

mv poj.org $id && cd $id && mv problem*$id ${id}.html

if [ $? -ne 0 ]; then
    echo $?
    exit 1
fi
