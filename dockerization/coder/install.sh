#!/bin/bash

curl -sSfL 'https://github.com/GaiaNet-AI/gaianet-node/releases/latest/download/install.sh' | bash

source ~/.bashrc

/root/gaianet/bin/gaianet init --config https://raw.githubusercontent.com/Gaianet-AI/node-configs/main/qwen-2.5-coder-0.5b-instruct/config.json

/root/gaianet/bin/gaianet config --domain gaia.domains

/root/gaianet/bin/gaianet init

/root/gaianet/bin/gaianet start

/root/gaianet/bin/gaianet info