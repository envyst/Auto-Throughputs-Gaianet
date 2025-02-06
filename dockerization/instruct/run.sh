#!/bin/bash

docker run -dit -P --restart unless-stopped gaianet-instruct 
docker ps