#!/bin/env bash
#
CMD="podman run --rm -it kafka:3.4 bash -c"

while getopts 'lcdsp' OPTION; do
  case "$OPTION" in
    l)
      echo "listing"
      ${CMD} "kafka-topics.sh --bootstrap-server kafka:9092 --list"
      ;;
    c)
      echo "creating $2"
      ${CMD} "kafka-topics.sh --bootstrap-server kafka:9092 --create --topic $2"
      ${CMD} "kafka-configs.sh --bootstrap-server kafka:9092 --alter --entity-type topics --entity-name $2 --add-config retention.ms=300000"
      ;;
    d)
      echo "deleting ${TOPIC}"
      ${CMD} "kafka-topics.sh --bootstrap-server kafka:9092 --delete --topic $2"
      ;;
    s)
      echo "consuming"
      ${CMD} "kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic $2"
      ;;
    p)
      echo "producing"
      ${CMD} "kafka-console-producer.sh --bootstrap-server kafka:9092 --topic $2"
      ;;
    ?)
      echo "script usage: $(basename \$0) [-l] [-h] [-a somevalue]" >&2
      exit 1
      ;;
  esac
done

shift "$(($OPTIND -1))"
