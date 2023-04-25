#!/bin/bash

input=`cat`
echo "check $input"

if [ "$input" -le "$1" ]; then
    echo "yes" >&3
else
    echo "no" >&3
fi
