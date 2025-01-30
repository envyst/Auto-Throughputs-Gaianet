#!/bin/bash

/root/gaianet/bin/gaianet init
/root/gaianet/bin/gaianet start

# Extract Node ID
node_id=$(jq -r '.address' nodeid.json)

# Extract Device ID
device_id=$(cat /root/gaianet/deviceid.txt)

# Print the variables to verify
echo "Node ID: $node_id"
echo "Device ID: $device_id"

# Write the variables to a file (optional)
echo "------------------------------------------------------------" >> /root/gaianet/node_info/node_info.txt
echo "Node ID: $node_id" >> /root/gaianet/node_info/node_info.txt
echo "Device ID: $device_id" >> /root/gaianet/node_info/node_info.txt
echo "------------------------------------------------------------" >> /root/gaianet/node_info/node_info.txt

echo "Startup complete"