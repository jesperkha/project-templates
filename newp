#!/usr/bin/env bash

# ----------------------------
#
#    SET TO THIS (REPO) DIRECTORY
# 
#    sed -i "s|^TEMPLATES_DIR=.*|TEMPLATES_DIR=$(pwd)|" newp
#
# ----------------------------
TEMPLATES_DIR=

if [[ -z "$TEMPLATES_DIR" ]]; then
    echo "ERROR: Template directory not set. Edit TEMPLATES_DIR in newp script."
    exit 1
fi

# Check that templates directory exists
if [[ ! -d "$TEMPLATES_DIR" ]]; then
    echo "ERROR: Templates directory not found: $TEMPLATES_DIR"
    exit 1
fi

# Prompt user for project type
echo "Available templates:"
for dir in "$TEMPLATES_DIR"/*/; do
    echo "  $(basename "$dir")"
done
echo

read -rp "Enter project type (template name): " TEMPLATE

# Check that the chosen template exists
if [[ ! -d "$TEMPLATES_DIR/$TEMPLATE" ]]; then
    echo "ERROR: Template '$TEMPLATE' does not exist."
    exit 1
fi

# Ask where to put the template
echo
echo "Where do you want to create the project?"
echo "1) Current directory ($(pwd))"
echo "2) New directory"
read -rp "Choose 1 or 2: " CHOICE

TARGET_DIR=""

if [[ "$CHOICE" == "1" ]]; then
    TARGET_DIR="$(pwd)"
elif [[ "$CHOICE" == "2" ]]; then
    read -rp "Enter name for the new directory: " NEW_DIR
    if [[ -z "$NEW_DIR" ]]; then
        echo "ERROR: Directory name cannot be empty."
        exit 1
    fi
    TARGET_DIR="$(pwd)/$NEW_DIR"
    mkdir -p "$TARGET_DIR"
else
    echo "ERROR: Invalid choice."
    exit 1
fi

# Copy template into target directory
echo "Copying template '$TEMPLATE' into '$TARGET_DIR'..."
cp -r "$TEMPLATES_DIR/$TEMPLATE/." "$TARGET_DIR"

echo "Done! '$TEMPLATE' template is ready in '$TARGET_DIR'."
