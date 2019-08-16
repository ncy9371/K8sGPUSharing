#!/bin/bash

function trap_ctrlc ()
{
    echo "Ctrl-C caught...performing clean up"
    for pid in $pids; do
        kill $pid
    done
    echo "Doing cleanup"
    exit 0
}

trap "trap_ctrlc" 2

port=50051

command -v nvidia-smi
if [ $? -ne 0 ]; then
    echo "No GPU available, sleep forever"
    sleep infinity
fi

for gpu in $(nvidia-smi --format=csv,noheader --query-gpu=uuid); do
    ./launcher.sh $port &
    pids="$pids $!"
    port=$(($port+1))
done

wait
