#!/bin/bash
#
METHOD=$(cat $1 | yq '.method')
REQUEST=$(cat $1 | yq '.request' -o=json)

grpc_cli call --json_input --json_output localhost:80 $METHOD "$REQUEST" | jq 
