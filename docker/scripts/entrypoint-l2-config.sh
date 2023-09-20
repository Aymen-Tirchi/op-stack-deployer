#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/6_L2_config/main.go

touch /app/L2_config_done