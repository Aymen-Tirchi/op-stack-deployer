#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/1_build_optimism/main.go

touch /app/build-optimism_done