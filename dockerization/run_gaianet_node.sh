#!/bin/bash

sudo chmod +x install_docker.sh
./install_docker.sh
docker compose up -d

echo "Waiting for the node to start..."
sleep 5

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
echo "------------------------------------------------------------"