#!/bin/bash

curl -sSfL 'https://github.com/GaiaNet-AI/gaianet-node/releases/latest/download/install.sh' | bash

source ~/.bashrc

/root/gaianet/bin/gaianet init --config https://raw.githubusercontent.com/Gaianet-AI/node-configs/main/qwen2-0.5b-instruct/config.json

/root/gaianet/bin/gaianet config --domain gaia.domains

/root/gaianet/bin/gaianet init

/root/gaianet/bin/gaianet start

/root/gaianet/bin/gaianet info