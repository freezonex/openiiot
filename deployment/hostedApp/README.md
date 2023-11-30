## Refered Document
https://ofra65wfwe.feishu.cn/wiki/PxgLwDBrQi87OrkFtOkcJa8FnJh
## Installation
The installation guide below is only suitable for SupOS private installation. For ESS/APaaS version, you can simply buy the app from the app store
### Docker Image Push
1. Use FTP or SCP to upload the docker image to SupOS server
```shell
$ scp node-red.tar username@remote_host:/path/to/destination
```
2. Ssh to SuoOS host and load the uploaded docker image
```shell
$ ssh root@http://47.236.10.165/ 
$ docker load -i node-red.tar
$ docker images
REPOSITORY                                 TAG              IMAGE ID       CREATED             SIZE
node-red-linux                             1.0.0            d68f19f85143   About an hour ago   1.03GB
```
3. In SupOS host, check Dockerhub image namespace, take 47 for example, password is Supos@1304
```shell
% docker image ls
REPOSITORY                                                  TAG                            IMAGE ID       CREATED         SIZE
registry:5000/dt/javaw-dname                                1.0.0                          5c35e6961913   6 days ago      126MB
registry.supos.ai/jenkins/lake-gateway                      4.02.01.29-C-M1-T2             ed64a2b8e5bd   7 weeks ago     54.6MB
registry:5000/jenkins/lake-gateway                          4.02.01.29-C-M1-T2             ed64a2b8e5bd   7 weeks ago     54.6MB
registry.supos.ai/jenkins/rbac                              4.02.01.29-C-M1-T4             4e43d4444a6b   7 weeks ago     557MB
registry:5000/jenkins/rbac                                  4.02.01.29-C-M1-T4             4e43d4444a6b   7 weeks ago     557MB
registry:5000/jenkins/websocketserver                       4.02.01.29-M1-T1               b782ba5d7282   7 weeks ago     412MB
registry.supos.ai/jenkins/websocketserver                   4.02.01.29-M1-T1               b782ba5d7282   7 weeks ago     412MB
registry.supos.ai/jenkins/dispatchor                        4.02.01.29-C-R1-T2             6e1d76b6ceac   7 weeks ago     492MB
registry:5000/jenkins/dispatchor                            4.02.01.29-C-R1-T2             6e1d76b6ceac   7 weeks ago     492MB
registry.supos.ai/jenkins/lake-rtdbroker                    4.02.01.29-C-M1-T1             5324cd5703a4   7 weeks ago     59.8MB
registry:5000/jenkins/lake-rtdbroker                        4.02.01.29-C-M1-T1             5324cd5703a4   7 weeks ago     59.8MB
registry.supos.ai/jenkins/supos-gateway                     4.2.1.29-T1                    6954fcf342f2   7 weeks ago     420MB
registry:5000/jenkins/supos-gateway                         4.2.1.29-T1                    6954fcf342f2   7 weeks ago     420MB
registry.supos.ai/jenkins/notification-integration          4.02.01.29-C-M1-T1             d6d095938f49   7 weeks ago     573MB
registry:5000/jenkins/notification-integration              4.02.01.29-C-M1-T1             d6d095938f49   7 weeks ago     573MB
registry.supos.ai/jenkins/oodm-frontend                     4.02.01.29-C-M1-T2             c3a9732b4568   7 weeks ago     33MB
```
You will find the dockerhub namespace is "registry:5000", then tag your docker image as 
```shell
docker tag tools-node-red:latest registry:5000/node-red:1.0.0
```
4. Push to SupOS dockerhub
```shell
docker push registry:5000/node-red:1.0.0
```
### Modify zip Package
All hosted app zip packages using 7zip, using other zip apps may cause installation packages to fail to load  
Windows is easy to modify zip by 7zip file manager and no need to extract  
For mac, use 7z console, download 7z console from official website: https://www.7-zip.org/download.html, tke 7z in mac for example
```shell
7zz x openiiot.zip
```
after modify yml file, then update back to openiiot.zip
```shell
7zz u openiiot.zip compose.yaml
```
