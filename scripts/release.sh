#!/bin/sh

rm -rf pkg/*
mkdir -p pkg

gox \
  -output "pkg/{{.OS}}_{{.Arch}}/aero" \
  -osarch "darwin/amd64" \
  -osarch "linux/amd64"

for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
  OSARCH=$(basename ${PLATFORM})
  pushd $PLATFORM >/dev/null 2>&1
  zip ../aero_${OSARCH}.zip ./*
  popd >/dev/null 2>&1
done
