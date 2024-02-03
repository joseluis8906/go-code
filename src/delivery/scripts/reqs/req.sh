#!/bin/bash
#
METHOD=$(cat $1 | yq '.method')
METADATA=$(cat $1 | yq '.metadata')
REQUEST=$(cat $1 | yq '.request' -o=json)

grpc_cli call --json_input --json_output localhost:50051 $METHOD "$REQUEST" --metadata="$METADATA" | jq 
