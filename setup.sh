#!/usr/bin/env bash

cp newp_original.sh newp.sh
chmod +x newp.sh

# Set TEMPLATES_DIR in the script to current directory
sed -i "s|^TEMPLATES_DIR=.*|TEMPLATES_DIR=$(pwd)|" newp.sh

# Ask where to install the script
DEFAULT_DIR="$HOME/.local/bin"
read -rp "Where do you want to put the script? [$DEFAULT_DIR]: " INSTALL_DIR
INSTALL_DIR="${INSTALL_DIR:-$DEFAULT_DIR}"

# Create target directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Create symlink in the chosen directory
# ln -sf "$(pwd)/newp.sh" "$INSTALL_DIR/newp"

# Copy script to install dir
cp newp.sh $INSTALL_DIR/newp

echo "Script installed as '$INSTALL_DIR/newp'. Make sure '$INSTALL_DIR' is in your PATH."
