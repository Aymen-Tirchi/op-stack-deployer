#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/4_configure_network/main.go

touch /app/configure_network_done