#!/usr/bin/env sh

result=0

while IFS= read -r line; do

    # handle overlaps
    set -- "one" "two" "three" "four" "five" "six" "seven" "eight" "nine"
    for w1 in "$@"; do
        c1=$(echo "$w1" | cut -c 1)
        for w2 in "$@"; do
            c2=$(echo "$w2" | rev | cut -c 1)
            if test "$c1" = "$c2"; then
                line="$(echo "$line" | sed "s/${w2}$(echo "$w1" | cut -c2-)/${w2}${w1}/g")"
            fi
        done
    done

    # replace words
    line=$(
        echo "$line" | sed \
            -e 's/one/1/g' \
            -e 's/two/2/g' \
            -e 's/three/3/g' \
            -e 's/four/4/g' \
            -e 's/five/5/g' \
            -e 's/six/6/g' \
            -e 's/seven/7/g' \
            -e 's/eight/8/g' \
            -e 's/nine/9/g'
    )

    digits="$(echo "${line}" | tr -d '[:alpha:]')"
    summand="$(echo "$digits" | cut -c 1)$(echo "$digits" | rev | cut -c 1)"
    result=$((result + summand))

done <"input.txt"

echo $result
