FROM ubuntu:22.04

RUN apt-get update \
    && apt-get install curl jq -y

RUN curl -sSfL 'https://github.com/GaiaNet-AI/gaianet-node/releases/latest/download/install.sh' | bash
RUN rm -rf /root/gaianet/config.json
RUN . ~/.bashrc 

COPY config.json /root/gaianet/config.json

RUN /root/gaianet/bin/gaianet init
CMD ["tail", "-f", "/dev/null"]
