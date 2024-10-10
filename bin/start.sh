#!/bin/bash

cp ./pls /usr/local/bin/pls
chmod +x /usr/local/bin/pls

echo "pls installed successfully"

nohup pls serve > log.log 2>&1 &
