#!/bin/bash

# server
cd apps/server/main/
go build -o $HOME/bin/server

cd ../../client/main/

# client
go build -o $HOME/bin/client

cd ../../../
