#!/bin/bash

set -e

APP_NAME="dm"
BUILD_OUTPUT="./$APP_NAME"
TARGET_PATH="/usr/local/bin/$APP_NAME"
SERVICE_NAME="dm"

echo "Building $APP_NAME..."
go build -o "$BUILD_OUTPUT" .

echo "Stopping service $SERVICE_NAME..."
sudo systemctl stop "$SERVICE_NAME"

echo "Deploying binary to $TARGET_PATH..."
sudo cp "$BUILD_OUTPUT" "$TARGET_PATH"
sudo chmod +x "$TARGET_PATH"

echo "Starting service $SERVICE_NAME..."
sudo systemctl start "$SERVICE_NAME"

echo "Deployment complete."
