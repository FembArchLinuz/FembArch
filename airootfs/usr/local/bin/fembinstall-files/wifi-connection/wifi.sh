#!/usr/bin/env bash

echo "[INFO]: starting iwd..."
iwctl
echo "[INFO]: Okay boom you connected to the internet i guess."
sleep 0.5
dhcpcd
