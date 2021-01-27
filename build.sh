#!/bin/bash
version=$1
releasebuild=1.0.0
commithash=$(git log –pretty=format:’%h’ -n 1)
extension=""
if [ "$GOOS" == "windows" ]
then
    extension=".exe"
elif [ "$GOOS" == "darwin" ]
then
    extension="_osx"
fi

filename="xrctournament_$version$extension"

if [ "$version" == "pro" ]
then
    if [ "$GOOS" == "linux" ]
    then
        go test -v --tags pro
    fi
    go build -v -tags pro -ldflags "-s -w -X main.Version=$releasebuild -X main.CommitHash=$commithash" -o "$filename"
else
    if [ "$GOOS" == "linux" ]
    then
        go test -v --tags free
    fi
    go build -v -tags free -ldflags "-s -w -X main.Version=$releasebuild -X main.CommitHash=$commithash" -o "$filename"
fi
