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
    go build -v -tags pro -o "$filename"
else
    go build -v -tags free -o "$filename"
fi
