#!/bin/bash

sudo chmod +x install_docker.sh
./install_docker.sh
docker compose up -d
# Wait for a specific log entry from the main container
CONTAINER_NAME="gaianet_node"  # Change this to your actual container name
SUCCESS_MESSAGE="Startup complete"    # Change this to the expected log message

echo "Waiting for $CONTAINER_NAME to complete startup..."

while true; do
    if docker logs "$CONTAINER_NAME" 2>&1 | grep -q "$SUCCESS_MESSAGE"; then
        echo "Docker entrypoint script finished!"
        break
    fi
    sleep 2  # Check every 2 seconds
done

# Run the next command
echo "Done"
# Run the command and capture the output
output=$(docker exec -it gaianet_node /root/gaianet/bin/gaianet info)

# Extract Node ID
node_id=$(echo "$output" | sed -n 's/Node ID: \(.*\)/\1/p')

# Extract Device ID
device_id=$(echo "$output" | sed -n 's/Device ID: \(.*\)/\1/p')

# Print the variables to verify
echo "Node ID: $node_id"
echo "Device ID: $device_id"

# auto throughput
screen -S gaianet-docker-throughput
cd ../go-app/
./program -nodeID $node_id