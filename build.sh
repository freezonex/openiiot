#!/bin/bash

RUN_NAME="freezonex.openiiot"

mkdir -p output/bin output/conf
cp biz/config/conf/*.yml output/conf/
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
cp script/bootstrap.sh output/bootstrap_staging.sh
chmod +x output/bootstrap_staging.sh
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

go build -o output/bin/${RUN_NAME}

# build modbus simulator
MODBUS_SERVER_NAME="modbus_server"
MODBUS_CLIENT_NAME="modbus_client"

mkdir -p output/bin

go build -o output/bin/${MODBUS_SERVER_NAME} simulator/modbus/server/main.go

go build -o output/bin/${MODBUS_CLIENT_NAME} simulator/modbus/client/main.go

