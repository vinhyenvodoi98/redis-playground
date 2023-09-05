#!/bin/bash

total_inserts=1000000

for ((i = 1; i <= total_inserts; i++)); do
    key=$(uuidgen)
    value=$(uuidgen)
    ttl=$((RANDOM % 36000000))

    redis-cli SETEX "$key" "$ttl" "$value"
done
