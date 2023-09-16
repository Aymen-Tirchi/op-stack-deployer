#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/3_generate_keys/main.go

touch /app/generate_keys_done