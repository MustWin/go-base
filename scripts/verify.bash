#!/bin/bash

# This is called from pre-push.bash to do some verification checks on
# the Go code. The script will exit non-zero if any of these tests fail.
# However if environment variable IGNORE_VET_WARNINGS is a non-zero
# length string, go vet warnings will not exit non-zero. Also, if
# IGNORE_TEST_ERRORS is non-empty then test failures will not exit
# non-zero either.

set -e 

src_folders="app base commands core test"
packages=$(for s in ${src_folders}; do echo -e "./$s/..."; done)

VERSION=$(go version | awk '{print $3}')
echo "go version $VERSION"

echo "checking: go fmt ..."
BADFMT=$(find $src_folders -name '*.go' | xargs gofmt -l)
if [ -n "$BADFMT" ]; then
    BADFMT=`echo "$BADFMT" | sed "s/^/  /"`
    echo -e "go fmt is sad:\n\n$BADFMT"
    exit 1
fi


echo "checking: go vet ..."
for dir in $src_folders; do echo " ... ${dir}"
	go tool vet \
			-methods \
			-printf \
			-rangeloops \
			${dir} || [ -n "$IGNORE_VET_WARNINGS" ]
done


echo "checking: go build ..."
# check this branch builds cleanly
go build $packages


echo "checking: go test ..."
go test $packages || [ -n "$IGNORE_TEST_ERRORS" ]
