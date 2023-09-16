#!/bin/sh

while [ ! -f /app/previous_step_done ]; do 
    sleep 1
done

go run cmd/5_deploy_L1_contracts/main.go

touch /app/deploy_L1_contracts_done