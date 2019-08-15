#!/bin/bash

MODE=$1
COMMIT=${@:2}

function buildBinaries {
    echo "-----> Building binaries"
    # build a binary in root directory for your current environment
    go build -o xkcd .
    # build all binaries in ./bin/
    for os in windows linux darwin
        do
        for arch in amd64 386
            do
            mkdir -p ./bin/$arch/$os/
            env GOOS=$os GOARCH=$arch go build -v -o ./bin/$arch/$os/xkcd .
            done
        done
}

function gitCommit {
    echo "-----> Committing changes to Git"
    git add -A
    git commit -m "$COMMIT"
}

buildBinaries
if [ "$MODE" == "commit" ]; then
    gitCommit
fi