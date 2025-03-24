#!/bin/sh
set -e

if ! getent group kleister >/dev/null 2>&1; then
    groupadd --system kleister
fi

if ! getent passwd kleister >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/kleister --shell /bin/bash -g kleister kleister
fi
