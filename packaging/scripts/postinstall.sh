#!/bin/sh
set -e

chown -R kleister:kleister /etc/kleister
chown -R kleister:kleister /var/lib/kleister
chmod 750 /var/lib/kleister

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet kleister-api.service; then
        systemctl restart kleister-api.service
    fi
fi
