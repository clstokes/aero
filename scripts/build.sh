#!/bin/sh

rm -rf bin/*
mkdir -p bin

gox \
  -output "bin/aero" \
  -osarch "$(go env GOOS)/$(go env GOARCH)"
