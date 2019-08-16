#!/bin/bash

function trap_myctrlc ()
{
	kill -9 $mypid
    exit 15
}

trap "trap_myctrlc" 15

while true; do
    ./scheduler $1 &
    mypid="$!"
    wait
    echo "ERROR! $1 crash!"
done
