#!/bin/bash

src_folders="app base commands test"
packages=$(for s in ${src_folders}; do echo -e "./$s/..."; done)

VERSION=$(go version | awk '{print $3}')
echo "go version $VERSION"

echo "checking: go lint ..."
for pkg in $packages; do echo " ... ${pkg}"
	golint -min_confidence 0.6 ${pkg}
done
