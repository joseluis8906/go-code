#!/usr/bin/env bash

SERVICE_NAME='sales_http'
SERVICE_PATH='sales_route'

c_flag=''
d_flag=''

while getopts cdv: flag
do
    case "${flag}" in
        c) c_flag='true' ;;
        d) d_flag='true' ;;
    esac
done

if [ "$c_flag" == 'true' ]; then
    curl -s -X POST http://localhost:8001/services \
        --data name=$SERVICE_NAME \
        --data url='http://sales-http:8080' | jq . && \

    curl -X POST http://localhost:8001/services/$SERVICE_NAME/routes \
        --data 'paths[]=/sales' \
        --data 'hosts[]=api.example.com' \
        --data name=$SERVICE_PATH | jq .
else
    curl -X DELETE http://localhost:8001/services/$SERVICE_NAME/routes/$SERVICE_PATH
    curl -X DELETE http://localhost:8001/services/$SERVICE_NAME
fi
