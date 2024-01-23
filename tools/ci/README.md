## Create a bridge network in Docker using the following docker network create command
```shell
docker network create jenkins
```
## Download and run the docker:dind Docker image
```shell
docker run \
  --name jenkins-docker \
  --rm \
  --detach \
  --privileged \
  --network jenkins \
  --network-alias docker \
  --env DOCKER_TLS_CERTDIR=/certs \
  --volume jenkins-docker-certs:/certs/client \
  --volume jenkins-data:/var/jenkins_home \
  --publish 2376:2376 \
  docker:dind \
  --storage-driver overlay2
```
## Build docker image
```shell
docker build -t myjenkins-blueocean:2.426.2-1 .
```
## Run Jenkins container
```shell
docker run \
  --name jenkins-blueocean \
  --restart=on-failure \
  --detach \
  --network jenkins \
  --env DOCKER_HOST=tcp://docker:2376 \
  --env DOCKER_CERT_PATH=/certs/client \
  --env DOCKER_TLS_VERIFY=1 \
  --publish 49000:8080 \
  --publish 50000:50000 \
  --volume jenkins-data:/var/jenkins_home \
  --volume jenkins-docker-certs:/certs/client:ro \
  myjenkins-blueocean:2.426.2-1
```

## Add a New Webhook in pb_idl project:

Payload URL: http://47.245.114.166:49000/github-webhook/
Content Type: application/json
Secret: Supos@1304
Event: Just push event

## Configure Jenkins Pipeline Job
| Option | Value |
| ------ | ----- |
| Name | pipeline_pb_idl | 
| Description | Triggers hz update command in openiiot when merge code to pb_idl main branch |
| Option | Do not allow concurrent builds |
| GitHub project | https://github.com/freezonex/pb_idl |
| Build Triggers | GitHub hook trigger for GITScm polling |
| Pipeline script | pipeline_pb_idl.groovy |

## Golang Agent
build golang agent use agent/Dockerfile
```shell
docker build -t jenkins-golang-agent .
```

start golang agent
```shell
docker run -d --name jenkins-golang-agent \
    -v /var/jenkins_agent:/var/jenkins_agent \
    -w /var/jenkins_agent \
    jenkins-golang-agent:latest \
    bash -c "curl -sO http://47.245.114.166:49000/jnlpJars/agent.jar && \
             java -jar agent.jar -jnlpUrl http://47.245.114.166:49000/computer/jenkins%2Dgolang%2Dagent/jenkins-agent.jnlp -secret 835fbbca9038f8555e7984a68001042c10ffa5b8f399633473ea167555f7024d -workDir '/var/jenkins_agent'"
```

if get connection error, run the container in interactive mode
```shell
docker run -it --name jenkins-golang-agent \
    -v /var/jenkins_agent:/var/jenkins_agent \
    -w /var/jenkins_agent \
    jenkins-golang-agent:latest bash
```