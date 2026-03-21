#!/bin/bash

# This script cleans the subdirectories by removing generated files.
# It targets mus_test.go and mus.ai.gen.go.

# Target directory is the first argument, or the script's directory if none provided.
TARGET_DIR="${1:-$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )}"

echo "Cleaning generated files in $TARGET_DIR..."

# Find and delete the specified files in subdirectories.
# Use -mindepth 1 to look into subdirectories of the target directory.
find "$TARGET_DIR" -mindepth 1 -type f \( -name "mus_test.go" -o -name "mus.ai.gen.go" \) -delete

echo "Done."
