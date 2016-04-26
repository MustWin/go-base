#!/bin/bash

src_folders="app base commands test"

# Perform import group checking and ordering
goimports -w ${src_folders}

# Perform simplification 
gofmt -s -w ${src_folders}
