#!/bin/sh

set -e

VERSION=$1
if [ -z "${VERSION}" ]; then
  echo "Usage: ${0} <version>" >> /dev/stderr
  exit 255
fi

rm -rf pkg/*
mkdir -p pkg

gox \
  -output "pkg/{{.OS}}_{{.Arch}}/aero" \
  -osarch "darwin/amd64" \
  -osarch "linux/amd64"

for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
  OSARCH=$(basename ${PLATFORM})
  pushd $PLATFORM >/dev/null 2>&1
  zip ../aero_${VERSION}_${OSARCH}.zip ./*
  popd >/dev/null 2>&1
done

pushd pkg >/dev/null 2>&1
shasum -a256 *.zip > aero_SHA256SUMS
popd >/dev/null 2>&1
