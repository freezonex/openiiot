#!/bin/bash

MODBUS_SERVER_NAME="modbus_server"
MODBUS_CLIENT_NAME="modbus_client"

mkdir -p output/bin

go build -o output/bin/${MODBUS_SERVER_NAME} simulator/modbus/server/main.go

go build -o output/bin/${MODBUS_CLIENT_NAME} simulator/modbus/client/main.go

