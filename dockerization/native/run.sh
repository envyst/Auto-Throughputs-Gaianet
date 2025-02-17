#!/bin/bash

docker run -dit -P --restart unless-stopped -v $(pwd)/qdrant_storage:/root/gaianet/qdrant/storage:z gaianet
docker ps