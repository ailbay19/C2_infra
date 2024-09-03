#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

FILENAME=$1

FILEPATH="./$FILENAME"

# Check if the file exists
if [ ! -f "$FILENAME" ]; then
    echo "File not found: $FILENAME"
    exit 1
fi

# Define the URL to post to
URL="https://localhost:443/upload"

curl -X POST -F "file=@$FILEPATH" --cert ./ssl/client.crt --key ./ssl/client.key --cacert ./ssl/server.crt "$URL"
