#!/bin/bash

/root/gaianet/bin/gaianet init
/root/gaianet/bin/gaianet start

# Run the command and capture the output
output=$(/root/gaianet/bin/gaianet info)

# Extract Node ID
node_id=$(echo "$output" | sed -n 's/Node ID: \(.*\)/\1/p')

# Extract Device ID
device_id=$(echo "$output" | sed -n 's/Device ID: \(.*\)/\1/p')

# Print the variables to verify
echo "Node ID: $node_id"
echo "Device ID: $device_id"

# Write the variables to a file (optional)
echo "------------------------------------------------------------" >> /root/gaianet/node_info/node_info.txt
echo "Node ID: $node_id" >> /root/gaianet/node_info/node_info.txt
echo "Device ID: $device_id" >> /root/gaianet/node_info/node_info.txt
echo "------------------------------------------------------------" >> /root/gaianet/node_info/node_info.txt

echo "Startup complete"