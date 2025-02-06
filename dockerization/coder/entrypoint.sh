#!/bin/bash

FLAG_FILE="/first_run_done"

if [ ! -f "$FLAG_FILE" ]; then
    echo "First time startup"
    /install.sh
    touch "$FLAG_FILE"  # Mark that the first run has been completed
else
    echo "Container restarted"
    /restart.sh
fi

exec /bin/bash