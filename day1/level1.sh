#!/usr/bin/env sh

set -eu

result=0

while IFS= read -r line; do
    digits="$(echo "${line}" | tr -d '[:alpha:]')"
    summand="$(echo "$digits" | cut -c 1)$(echo "$digits" | rev | cut -c 1)"
    result=$((result + summand))
done <"input.txt"

echo $result
