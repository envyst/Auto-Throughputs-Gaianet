FROM ubuntu:latest

RUN apt-get update \
    && apt-get upgrade -y

RUN apt-get install -y \
    curl nano lsof jq screen

COPY install.sh /install.sh
COPY restart.sh /restart.sh
COPY entrypoint.sh /entrypoint.sh 

RUN chmod +x /install.sh /restart.sh /entrypoint.sh

EXPOSE 8080

WORKDIR /truput

COPY agt .
COPY api.txt .
COPY cersex.txt .

ENTRYPOINT ["/entrypoint.sh"]