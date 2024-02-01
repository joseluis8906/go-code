#!/bin/bash
#
# Set the IFS (Internal Field Separator) to space
IFS='/'
# Use read to bash convert string to array
read -ra strarr <<< "$1"
IFS=' '

len=${#strarr[@]}
PKG="${strarr[len-1]}"

dlv test --headless --listen=:2345 --api-version=2 github.com/joseluis8906/go-code/$1 -- -test.run ^Test$2
