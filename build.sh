#!/bin/bash
version=$1
extension=""
if [ "$GOOS" == "windows" ]
then
    extension=".exe"
elif [ "$GOOS" == "darwin" ]
    extension="_osx"
fi

filename="xrctournament_$version$extension"

if [ "$version" == "pro" ]
then
    go build -tags pro -v -ldflags "-s -w" -o $filename
else
    go build -v -ldflags "-s -w" -o $filename
fi