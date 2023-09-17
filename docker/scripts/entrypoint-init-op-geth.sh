#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/7_Initialize_op-geth/main.go

touch /app/Initialize_op-geth_done