#!/usr/bin/env bash

link_flags="-s -w"
comp_flags="-trimpath"
test_flags=-test.v

src_dir=./src
app_fname=build/pw
test_fname=test

res_origin="res"
res_destin="build/res"

go test -c -o $test_fname $src_dir
./$test_fname $test_flags

if [ $? == 0 ]
then
    go build $comp_flags -ldflags="${link_flags}" -o $app_fname $src_dir
    cp -R $res_origin $res_destin
    rm $test_fname
fi
