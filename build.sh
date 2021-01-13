#!/bin/bash
version=$1
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
    go test -v --tags pro
    go build -v -tags pro -ldflags "-s -w" -o "$filename"
else
    go test -v --tags free
    go build -v -tags free -ldflags "-s -w" -o "$filename"
fi
