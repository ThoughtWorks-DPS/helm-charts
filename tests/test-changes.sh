#!/usr/bin/env bash

# find all changed charts in the triggering commit
ct list-changed > charts-to-test

while read chartpath; do
  chart=${chartpath##*/}
  echo "testing $chart"
  cd ./tests/$chart
  go mod init "github.com/ThoughtWorks-DPS/helm-charts"
  go mod tidy
  go test -v -timeout 20m
  cd ../..
done <charts-to-test