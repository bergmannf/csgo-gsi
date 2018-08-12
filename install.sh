#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

CSGO_CFG_PATH="$HOME/.steam/steam/steamapps/common/Counter-Strike Global Offensive/csgo/cfg/"
CFG_NAME="gamestate_integration_sample.cfg"
CFG_SOURCE_PATH="./gsiConfigs/$CFG_NAME"
CFG_PATH="$CSGO_CFG_PATH/$CFG_NAME"

function install-config() {
    if [ ! -d "$CSGO_CFG_PATH" ]; then
        echo "$CSGO_CFG_PATH does not exists, will not copy config."
        exit 1
    fi
    cp "$CFG_SOURCE_PATH" "$CFG_PATH"
}

function uninstall-config() {
    if [ ! -f "$CFG_PATH" ]; then
        echo "$CFG_PATH not found. Nothing to uninstall."
        exit 1
    fi
    rm "$CFG_PATH"
}

while [ "$#" -gt 0 ]; do
    case "$1" in
        -i) install-config; shift 2;;
        -u) uninstall-config; shift 2;;

        -*) echo "unknown option: $1" >&2; exit 1;;
    esac
done
