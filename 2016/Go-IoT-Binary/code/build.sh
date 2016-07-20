#!/usr/bin/bash

echo "Checking GOPATH..."
echo ${GOPATH:?"-> not set, exiting..."}

cd $GOPATH
export GOOS=linux
export GOARCH=386

go build -v github.com/brockwood/govantage > build.log 2>&1

if [ "$?" -ne "0" ]; then
	echo "Error building the executable, please check the build log..."
else
	echo "Build was successful!"
fi