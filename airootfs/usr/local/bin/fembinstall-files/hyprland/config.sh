#!/usr/bin/env bash

# Make sure the config files exist inside of the .config/ directory.
# If .config/hypr doesn't exist then create it
if [[ ! -d "$HOME"/.config/hypr/ ]]; then
    mkdir -p "$HOME"/.config/hypr
elif [[ ! -f "$HOME"/.config/hypr/hyprland.conf ]]; then
    cp /usr/share/hypr/hyprland.conf "$HOME"/.config/hypr/hyprland.conf
fi

