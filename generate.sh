#!/usr/bin/env bash

mkdir "y$1"
cp -n template "y$1/$2.go"

sed -i "s|YYYY|y$1|g" "y$1/$2.go"
sed -i "s|DD|$2|g" "y$1/$2.go"
