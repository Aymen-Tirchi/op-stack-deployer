#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/12_get_rollup_address/main.go

touch /app/get_rollup_address_done