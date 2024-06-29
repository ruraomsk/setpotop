#!/bin/bash
echo "Start to Windows deploy"
GOOS=windows GOARCH=amd64 go build
if [ $? -ne 0 ]; then
	echo 'An error has occurred! Aborting the script execution...'
	exit 1
fi
echo "Start to Linux deploy"
go build
if [ $? -ne 0 ]; then
	echo 'An error has occurred! Aborting the script execution...'
	exit 1
fi
