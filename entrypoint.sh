#!/bin/sh

HOME=/srv/xlive

chown -R ${PUID}:${PGID} ${HOME}

umask ${UMASK}

exec su-exec ${PUID}:${PGID} /usr/bin/xlive -c /etc/xlive/config.yml
