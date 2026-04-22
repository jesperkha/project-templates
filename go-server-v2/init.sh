#!/usr/bin/env bash

OLD_NAME="go-server-v2"

# Prompt for new project name
read -rp "Enter new project name: " NEW_NAME

# Check if empty
if [[ -z "$NEW_NAME" ]]; then
  echo "Project name cannot be empty."
  exit 1
fi

# Confirm replacement
echo "Replacing all occurrences of '$OLD_NAME' with '$NEW_NAME'..."
read -rp "Continue? (y/n): " CONFIRM
if [[ "$CONFIRM" != "y" && "$CONFIRM" != "Y" ]]; then
  echo "Aborted."
  exit 0
fi

# Find and replace
# Excludes .git directory to avoid corrupting repository history
grep -rl --exclude-dir=.git "$OLD_NAME" . | while read -r file; do
  sed -i "s/$OLD_NAME/${NEW_NAME}/g" "$file"
done

echo "Replacement complete."
