#!/bin/env bash
#
CMD="ssh jose@10.0.0.2 podman exec -it common-infra_ejabberd_1 "

while getopts 'eru' OPTION; do
  case "$OPTION" in
    e)
      echo "status"
      ${CMD} "bin/ejabberdctl status"
      ;;
    r)
      echo "registering $2"
      ${CMD} "bin/ejabberdctl register $2 localhost $3"
      ;;
    u)
      echo "unregistering $2"
      ${CMD} "bin/ejabberdctl unregister $2 localhost"
      ;;
    ?)
      echo "script usage: $(basename \$0) [-l] [-h] [-a somevalue]" >&2
      exit 1
      ;;
  esac
done

shift "$(($OPTIND -1))"
