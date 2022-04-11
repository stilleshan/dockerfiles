#!/bin/sh

set -e

CONFIG_YML=/etc/docker/registry/config.yml

if [ -n "$PROXY_REMOTE_URL" -a `grep -c "$PROXY_REMOTE_URL" $CONFIG_YML` -eq 0 ]; then
    echo "proxy:" >> $CONFIG_YML
    echo "  remoteurl: $PROXY_REMOTE_URL" >> $CONFIG_YML
    echo "------ Enabled proxy to remote: $PROXY_REMOTE_URL ------"
elif [ $DELETE_ENABLED = true -a `grep -c "delete:" $CONFIG_YML` -eq 0 ]; then
    sed -i '/rootdirectory/a\  delete:' $CONFIG_YML
    sed -i '/delete/a\    enabled: true' $CONFIG_YML
    echo "------ Enabled local storage delete -----"
fi

case "$1" in
    *.yaml|*.yml) set -- registry serve "$@" ;;
    serve|garbage-collect|help|-*) set -- registry "$@" ;;
esac

exec "$@"
