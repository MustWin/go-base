#!/bin/bash

TARGET_DIRECTORY=${1:?first parameter should be target directory}
TARGET_IMPORTPATH=${2:?second parameter should be target import path}

mkdir -p ${TARGET_DIRECTORY}

git archive --format=tar HEAD | tar -C ${TARGET_DIRECTORY} -xf -

pushd ${TARGET_DIRECTORY}
git init
git add -A

find . -type f ! -path '*/vendor/*' ! -path '*/.git/*' | \
	xargs sed -i -e "s|github.com/MustWin/go-base|${TARGET_IMPORTPATH}|g"

popd

echo -e "\n${TARGET_DIRECTORY} has been populated with the files at the\n" \
"HEAD revision of this repository.  These files have been added to the\n" \
"git index, and staged.  Also, the import paths have been updated but\n" \
"not yet staged.\n\nNext step is to go to ${TARGET_DIRECTORY} and inspect/modify/commit.\n"
