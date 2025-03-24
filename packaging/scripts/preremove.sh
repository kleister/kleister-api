#!/bin/sh
set -e

systemctl stop kleister-api.service || true
systemctl disable kleister-api.service || true
