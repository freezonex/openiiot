## Development
### If develop in mainland China, need add use domestic mirror:
```shell
$ go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
$ go env | grep GOPROXY
```

### IDL
IDL repo: https://github.com/freezonex/pb_idl.git

Path: freezonex/openiiot_api/openiiot_api_service.proto

run below commands to pull the IDL repo automatically to the openiiot_api repo root folder.

```shell
git submodule add https://github.com/freezonex/pb_idl.git pb_idl
git submodule add https://github.com/freezonex/Freeflow-frontend.git frontend
```

```shell
git submodule init
git submodule update
```

### hertztool
we use hertz framework as the http service and it can generate the codes.
all the api service are defined in the IDL, it will generate the codes under biz/model.
when we update the IDL, we should raise MR to main branch in freezonex/pb_idl repo. and then it will trigger a background job to generate the code.
and push to openiiot main branch. what we need is to rebase from main branch.

refer here for detail.
#### installation of hertztool and protobuf support
```shell 
go install github.com/cloudwego/hertz/cmd/hz@latest
hz --version
brew install protobuf
protoc --version
```

if we want to test the IDL locally before merge to master branch, we can run below commands.
```shell 
hz model --idl pb_idl/freezonex/openiiot_api/openiiot_api_service.proto --unset_omitempty -I pb_idl
```

### DB
we are planning to connect to the openiiot_db, during the devleopment we can use docker to bring up local DB for the development and testing.

we use ORM library (gorm/GEN), and this library will be able to generate the basic code and function. note that, the configs and scripts are customized to generate the DB model/query. the generated codes are biz/dal/model and biz/data/query.

everytime we run the scripts, it will overwrite the codes. 
so far, the function JSONContains is a customized function to query a json column which contains a list of the tags.
```shell
sh generate.sh openiiot
```

### Configs
it use the yaml to do the configurations
details can be seen here. biz/config/*.yml
consts are defined under biz/config/consts/

### Build
build all docker images, default PLATFORM is linux/arm64
```shell
make build
```
if you want to build for mac:
```shell
make build PLATFORM=linux/arm64
```
for windows
```shell
make build PLATFORM=windows/amd64
```
build single image
```shell
make nodered
make grafana
make tdengine
make emqx
make mysql
make web
make server
make consolemanager
```
stop and delete all docker containers, images
```shell
make clean
```
stop and delete single docker containers, images
```shell
make clean_nodered
make clean_grafana
make clean_tdengine
make clean_emqx
make clean_mysql
make clean_web
make clean_server
make clean_consolemanager
```
save all images to deployment/binary
```shell
make save
```

save single image to deployment/binary
```shell
make save_nodered
make save_grafana
make save_tdengine
make save_emqx
make save_mysql
make save_web
make save_wever
make save_consolvemanager
```

### Push And Install
for Supos private installation, after upload all image to the server, just use script in deployment/install.sh to load, tag and push for all images
```shell
$ chmod +x install.sh
$ ./install.sh 
Loading image from openiiot_emqx.tar...
8ce178ff9f34: Loading layer [==================================================>]  84.03MB/84.03MB
ab3a1c678a1d: Loading layer [==================================================>]  7.988MB/7.988MB
ac8e1488e0fe: Loading layer [==================================================>]  346.6kB/346.6kB
88a076ed77e2: Loading layer [==================================================>]  142.1MB/142.1MB
0bfc0b9e58c7: Loading layer [==================================================>]  4.608kB/4.608kB
Loaded image: openiiot_emqx:1.0.0
Tagging image openiiot_emqx:1.0.0 as registry:5000/openiiot_emqx:1.0.0...
Pushing image registry:5000/openiiot_emqx:1.0.0...
The push refers to repository [registry:5000/openiiot_emqx]
0bfc0b9e58c7: Pushed 
88a076ed77e2: Pushed 
ac8e1488e0fe: Pushed 
ab3a1c678a1d: Pushed
```

### Unstall
just use script in deployment/uninstall.sh to remove all images
```shell
$ chmod +x uninstall.sh
$ ./uninstall.sh 
```

Project Objective:
To develop a comprehensive industrial platform designed for developers to deploy and manage industry-specific applications. The platform will integrate various IoT protocols to enable seamless connectivity with industrial devices, providing robust data management capabilities through advanced ETL processes and a dedicated data warehouse. It will empower users with tools to visualize data through custom dashboards and HMI (Human-Machine Interface) graphics. With built-in AI capabilities, the platform will offer predictive analytics and act as an AI assistant to enhance operational efficiency. User and tenant management features will ensure multi-level access control and resource management across different organizational layers.

Project Scope:
1. Platform Core:
Provide a unified platform for developers to install and manage various industrial applications.
Integrate with multiple IoT protocols (e.g., MQTT, OPC-UA, Modbus) to connect diverse industrial equipment and sensors.
Facilitate real-time data acquisition and monitoring from connected devices.

2. Data Management:
Include an ETL (Extract, Transform, Load) pipeline to process raw data from multiple sources into structured formats.
Establish a centralized data warehouse for storing historical and real-time data.
Enable seamless access to organized data for advanced analytics, reporting, and visualization.

3. Visualization and HMI:
Offer a customizable dashboard tool for developers to create real-time dashboards, reports, and data visualizations.
Provide HMI (Human-Machine Interface) design tools for graphical representations of industrial processes.
Support interactive visualizations for monitoring and controlling industrial operations.

4. AI Capabilities:
Implement AI-powered analytics to predict time-series data for predictive maintenance and operational insights.
Integrate an AI assistant to provide automated recommendations and decision support based on data patterns.

5. User and Tenant Management:
Support multi-tenancy with secure user management features to allow role-based access control.
Provide tenant isolation and resource allocation to ensure effective management of users across different organizational structures.

6. Developer-Focused Tools:
Deliver a suite of APIs and SDKs for developers to customize and extend platform functionality.
Facilitate integration with third-party industrial and IoT applications.
Ensure compatibility with popular development tools and frameworks to accelerate deployment.

7. Scalability and Security:
Ensure the platform is scalable to handle large datasets and numerous IoT connections.
Implement robust security measures for data protection, secure communication, and access control across all components.