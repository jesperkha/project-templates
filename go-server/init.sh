#!/usr/bin/env bash

# Prompt for new project name
read -rp "Enter new project name: " NEW_NAME

# Check if empty
if [[ -z "$NEW_NAME" ]]; then
  echo "Project name cannot be empty."
  exit 1
fi

# Confirm replacement
echo "Replacing all occurrences of 'go-server-template' with '$NEW_NAME'..."
read -rp "Continue? (y/n): " CONFIRM
if [[ "$CONFIRM" != "y" && "$CONFIRM" != "Y" ]]; then
  echo "Aborted."
  exit 0
fi

# Find and replace
# Excludes .git directory to avoid corrupting repository history
grep -rl --exclude-dir=.git "go-server-template" . | while read -r file; do
  sed -i "s/go-server-template/${NEW_NAME}/g" "$file"
done

sed -i "s/go-app/${NEW_NAME}/g" docker-compose.yaml

echo "Replacement complete."
