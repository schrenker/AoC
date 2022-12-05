#!/usr/bin/env sh

days=$(grep -A 9999 'var' challenge.go | tail -n +2 | sed  -e '$ d')

for i in $days; do
    tmp=$(echo "$i" | grep '/' | tr -d "\"" | tr -d ":")
    y=$(echo $tmp | cut -d/ -f1)
    d=$(echo $tmp | cut -d/ -f2)

    if [ -n "$y" ] && [ -n "$d" ]; then
    task run -- "$y" "$d" 01
    task run -- "$y" "$d" 02
    fi
done
