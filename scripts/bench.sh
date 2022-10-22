#!/usr/bin/env sh

YEAR=$1
DAY=$2

if [ "$3" = '1' ] || [ "$3" = '01' ]; then
    PART="One"
elif [ "$3" = '2' ] || [ "$3" = '02' ]; then
    PART="Two"
fi

go test -bench "Day${DAY}Part${PART}" ./benchmark/y"$YEAR"/...
