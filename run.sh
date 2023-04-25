#!/bin/bash

function print_use {
    exit 1
}

function parse_parameters {
    key=''
    cur=''
    checker_args=()
    simu_args=()
    args=()
    for arg in $@; do
        case $arg in
            '--checker') key='checker';;
            '--simu') key='simu';;
            '-x') key='env';;
            '--') key=''; cur='';;
            *)
            if [ "$key" == "" ]; then
                if [ "$cur" == 'simu' ]; then
                    simu_args+=("$arg")
                elif [ "$cur" == 'checker' ]; then
                    checker_args+=("$arg")
                else 
                    args+=("$arg")
                fi
            elif [ "$key" == 'env' ]; then
                export $arg
                key=''
            else
                export $key="$arg"
                cur=$key
                key=''
            fi
        esac
    done

    if [ "$checker" == "" ]; then
        print_use
    fi

    if [ "$simu" == "" ]; then
        print_use
    fi
}

parse_parameters "$@"

function clear {
    exec 3>&-

    rm pipes/input
    rm pipes/output
}

mkdir -p pipes
mkfifo pipes/input
mkfifo pipes/output

trap clear EXIT

$simu "${simu_args[@]}" < pipes/input 3> pipes/output &
pid=$!

exec 3> pipes/input
while read line; do
    if [[ "$line" == "end "* ]]; then
        echo ${line:4}
        break
    elif [[ "$line" == "check "* ]]; then
        $checker "${checker_args[@]}" < <(echo "${line:6}") 3> pipes/input
    fi
done < pipes/output
