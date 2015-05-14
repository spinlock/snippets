#!/bin/bash

init() {
    id=$1
    rm -rf $id
    mkdir -p $id && cd $id && wget -p -k http://poj.org/problem?id=$id

    if [ $? -ne 0 ]; then
        return $?
    fi

cat << EOF > Makefile
all:
	cat input | go run main.go

cc:
	clang++ -O2 code.cc -o /tmp/pojcode.out && cat input | /tmp/pojcode.out
EOF

    mv poj.org $id && cd $id && mv problem*$id ${id}.html

    if [ $? -ne 0 ]; then
        return $?
    fi

    cd ../../
    return 0
}

while true; do
    id=$1; shift

    if [ "x$id" == "x" ]; then
        exit 0
    fi

    init $id
    if [ $? -ne 0 ]; then
        printf "\e[1;31m[ERROR] %s %d\n\e[0m" $id $?
    else
        printf "\e[1;32m[OK] %s\n\e[0m" $id
    fi
done
