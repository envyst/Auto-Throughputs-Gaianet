#!/bin/bash

/root/gaianet/bin/gaianet config --domain gaia.domains
/root/gaianet/bin/gaianet init
/root/gaianet/bin/gaianet start
/root/gaianet/bin/gaianet info
if [ -f "~/api.txt" ]; then
    # If ~/api.txt exists, run ~/agt
    ~/agt
else
    # Otherwise, tail the log file
    tail -f /root/gaianet/log/start-llamaedge.log
fi