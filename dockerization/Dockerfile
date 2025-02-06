FROM ubuntu:latest

RUN apt-get update \
    && apt-get upgrade -y

RUN apt-get install -y \
    curl nano lsof jq

RUN curl -sSfL 'https://github.com/GaiaNet-AI/gaianet-node/releases/latest/download/install.sh' | bash

RUN rm -rf /root/gaianet/nodeid.json
RUN . ~/.bashrc 

COPY nodeid.json /root/gaianet/nodeid.json
COPY deviceid.txt /root/gaianet/deviceid.txt

EXPOSE 8080