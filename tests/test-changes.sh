#!/bin/bash

# find all changed charts in the triggering commit
ct list-changed > charts-to-test

echo "charts to test"
cat charts-to-test

if [ ! -s $file ]; then
  echo "no changed charts found"
fi
# while read chartpath; do
#   chart=${chartpath##*/}
#   echo ${#chart}
#   echo "testing $chart"
#   cd tests
#   cd $chart
#   pwd
#   cd ../..
#   pwd
#   # go mod init "github.com/ThoughtWorks-DPS/helm-charts"
#   # go mod tidy
#   # go test -v -timeout 20m
#   # cd ../..
# done <charts-to-test
