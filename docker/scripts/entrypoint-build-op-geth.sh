#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/2_build_op-geth/main.go

touch /app/build_op-geth_done