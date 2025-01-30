FROM ubuntu:22.04

# Install dependencies and setup Gaianet
RUN apt-get update && apt-get install -y curl && \
    curl -sSfL 'https://github.com/GaiaNet-AI/gaianet-node/releases/latest/download/install.sh' | bash && \
    rm -rf /root/gaianet/config.json

# Copy the config file
COPY config.json /root/gaianet/config.json

# Ensure .bashrc is sourced
RUN . ~/.bashrc

# Entrypoint script
COPY docker_entrypoint.sh /usr/local/bin/docker_entrypoint.sh
RUN chmod +x /usr/local/bin/docker_entrypoint.sh

# Expose port 3011
EXPOSE 3011

ENTRYPOINT ["docker_entrypoint.sh"]