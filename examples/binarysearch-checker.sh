#!/bin/bash

input=`cat`

if [ "$input" -le "$1" ]; then
    echo "yes" >&3
    echo "check $input return yes"
else
    echo "no" >&3
    echo "check $input return no"
fi
