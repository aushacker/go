#!/bin/bash

go build flags.go

echo "./flags -word=opt -numb=7 -fork -svar=flag"
./flags -word=opt -numb=7 -fork -svar=flag

echo
echo "./flags -word=opt"
./flags -word=opt

echo
echo "./flags -word=opt a1 a2 a3"
./flags -word=opt a1 a2 a3

echo
echo "./flags -word=opt a1 a2 a3-numb=7"
./flags -word=opt a1 a2 a3 -numb=7

echo
echo "./flags -h"
./flags -h

echo
echo "./flags -wat"
./flags -wat
