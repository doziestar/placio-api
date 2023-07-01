#!/bin/bash

# Set the script to exit immediately if any command fails
set -e

# Define variables for the host and port
host="localhost"
port="7070"

# Download wait-for-it.sh if it doesn't exist
if [ ! -f wait-for-it.sh ]; then
  echo "Downloading wait-for-it.sh..."
  curl -O https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
  chmod +x wait-for-it.sh
fi

# Start the server (replace this line with the actual command to start your server)
echo "Starting the server..."
docker-compose up --build -d

# Wait for the server to be available
echo "Waiting for the server to be available..."
bash ./wait-for-it.sh "${host}:${port}"

echo "Server is up - running tests..."

# Run the tests inside Docker container (replace "app" with your service name)
docker-compose exec -T app go test -v -coverprofile cover.out ./cmd/app/tests/...

# Generate the coverage report
echo "Generating the coverage report..."
docker-compose exec -T app go tool cover -html=cover.out -o cover.html

# Stop the server
echo "Stopping the server..."
docker-compose down

echo "Tests completed."
