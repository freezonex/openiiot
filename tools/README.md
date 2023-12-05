## Quick Start
execute in project root directory, this command will use existing docker images
```shell
docker-compose -f tools/docker-compose.yml up -d --no-build
```
execute in project root directory, this command will build docker images
```shell
docker-compose -f tools/docker-compose-build.yml up -d
```
stop and remove all containers:
```shell
docker-compose -f tools/docker-compose.yml down
```
stop single container only:
```shell
docker-compose -f tools/docker-compose.yml stop openiiot-nodered
```