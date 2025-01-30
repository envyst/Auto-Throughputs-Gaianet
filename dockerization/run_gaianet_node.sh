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

# Extract Node ID
node_id=$(docker exec gaianet_node jq -r '.address' nodeid.json)

# Extract Device ID
device_id=$(docker exec gaianet_node cat /root/gaianet/deviceid.txt)

# Print the variables to verify
echo "Node ID: $node_id"
echo "Device ID: $device_id"

echo "------------------------------------------------------------"
echo "Next run this command to start the throughput test:"
echo "screen -S gaianet-node-throughput"
echo "cd ../go-app/"
echo "./program -nodeID $node_id"
echo "Ctrl + A + D to detach from the screen session"