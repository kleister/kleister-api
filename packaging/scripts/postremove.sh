#!/bin/sh
set -e

if [ ! -d /var/lib/kleister ] && [ ! -d /etc/kleister ]; then
    userdel kleister 2>/dev/null || true
    groupdel kleister 2>/dev/null || true
fi
