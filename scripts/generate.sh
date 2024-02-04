#!/usr/bin/env sh

mkdir -p "y$1"
cp -n templates/template "y$1/$2.go"
sed -i "s|YYYY|y$1|g" "y$1/$2.go"
sed -i "s|DD|$2|g" "y$1/$2.go"

mkdir -p "benchmark/y$1"
cp -n templates/benchmark_template "benchmark/y$1/$2_test.go"
sed -i "s|YYYY|$1|g" "benchmark/y$1/$2_test.go"
sed -i "s|DD|$2|g" "benchmark/y$1/$2_test.go"

mkdir -p "test/y$1"
cp -n templates/test_template "test/y$1/$2_test.go"
sed -i "s|YYYY|$1|g" "test/y$1/$2_test.go"
sed -i "s|DD|$2|g" "test/y$1/$2_test.go"

mkdir -p "test/testdata/y$1/$2"
touch "test/testdata/y$1/$2/1.1.txt"
touch "test/testdata/y$1/$2/2.1.txt"

grep "$1/$2" challenge.go > /dev/null || tac challenge.go | sed -e "2i\"$1/$2\": y$1.Day$2{}," | tac | tee challenge.go > /dev/null

gofmt -w challenge.go
