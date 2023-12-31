## Quick Start
To use this repository, you need to install docker beforehand. if you get the message 'WSL kernel version too low' after installing docker desktop on Windows, run on the command line to upgrade it.
```shell
wsl --update
```
bring up a local TDengine DB, build and start the service
```shell
cd tools
docker-compose up -d 
```
## Configuration
### 1. Create TDengine Database and Table
In this step we create the appropriate database and table schema in TDengine for receiving MQTT data. 
First open Docker Terminal
![TDengine terminal](docs/images/TDengine_terminal.png "TDengine terminal")
Then create database by TDengine terminal:

```shell
# taos
taos>
```
Copy and paste below SQL script to taos
```SQL
CREATE DATABASE test;
USE test;
CREATE TABLE sensor_data (ts TIMESTAMP, temperature FLOAT, humidity FLOAT, volume FLOAT, pm10 FLOAT, pm25 FLOAT, so2 FLOAT, no2 FLOAT, co FLOAT, sensor_id NCHAR(255), area TINYINT, coll_time TIMESTAMP);
```
```SQL
CREATE DATABASE demo;
USE demo;
CREATE STABLE air_quality (
    ts TIMESTAMP,
    temperature FLOAT,
    humidity FLOAT,
    volume FLOAT,
    pm10 FLOAT,
    pm25 FLOAT,
    so2 FLOAT,
    no2 FLOAT,
    co FLOAT,
    status NCHAR(255))
    TAGS (node_id NCHAR(255), area NCHAR(255));
CREATE TABLE air_quality_1 USING air_quality TAGS ("node1", "Riyadh");
CREATE TABLE air_quality_2 USING air_quality TAGS ("node2", "Riyadh");
CREATE TABLE air_quality_3 USING air_quality TAGS ("node3", "Riyadh");
CREATE TABLE air_quality_4 USING air_quality TAGS ("node4", "Riyadh");
CREATE TABLE air_quality_5 USING air_quality TAGS ("node5", "AL-Madinah");
CREATE TABLE air_quality_6 USING air_quality TAGS ("node6", "AL-Madinah");
CREATE TABLE air_quality_7 USING air_quality TAGS ("node7", "AL-Madinah");
CREATE TABLE air_quality_8 USING air_quality TAGS ("node8", "Jeddah");
CREATE TABLE air_quality_9 USING air_quality TAGS ("node9", "Jeddah");
CREATE TABLE air_quality_10 USING air_quality TAGS ("node10", "Jeddah");
CREATE STABLE water_meter (
    ts TIMESTAMP,
    water_usage FLOAT,
    vibration_data FLOAT,
    status NCHAR(255))
    TAGS (node_id NCHAR(255), area NCHAR(255));
CREATE TABLE water_meter_1 USING water_meter TAGS ("node1", "Riyadh");
CREATE TABLE water_meter_2 USING water_meter TAGS ("node2", "Riyadh");
CREATE TABLE water_meter_3 USING water_meter TAGS ("node3", "Riyadh");
CREATE TABLE water_meter_4 USING water_meter TAGS ("node4", "Riyadh");
CREATE TABLE water_meter_5 USING water_meter TAGS ("node5", "AL-Madinah");
CREATE TABLE water_meter_6 USING water_meter TAGS ("node6", "AL-Madinah");
CREATE TABLE water_meter_7 USING water_meter TAGS ("node7", "AL-Madinah");
CREATE TABLE water_meter_8 USING water_meter TAGS ("node8", "Jeddah");
CREATE TABLE water_meter_9 USING water_meter TAGS ("node9", "Jeddah");
CREATE TABLE water_meter_10 USING water_meter TAGS ("node10", "Jeddah");
```
### 2. Create EMQX Integration to TDengine Rule
#### 2.1 Login EMQX Dashboard
Use your browser to open the URL http://IP:18083 (remember to change your IP here, and cannot use localhost or 127.0.0.1) and log in to EMQX Dashboard. Initial installation 
username: admin
password: public
![EMQX Login](docs/images/EMQX_login.png "EMQX Login")
#### 2.2 Create Data Bridges
Choose "Data Bridges" -> Click "Create" -> Choose "HTTP Servers" -> Configuration:
![EMQX data_bridges](docs/images/EMQX_data_bridges.png "EMQX data bridges")
in next page, enter below information:  
```http request
Name: Your name of the bridge, for example, "toTDengine"  
Method: POST  
URL: http://IP:9041/rest/sql, remember to change your IP here
Key: Authorization          Value: Basic cm9vdDp0YW9zZGF0YQ==
Query Mode: Async
```
Body Enter:
```SQL
INSERT INTO test.sensor_data VALUES(
  now,
  ${payload.temperature},
  ${payload.humidity},
  ${payload.volume},
  ${payload.PM10},
  ${payload.pm25},
  ${payload.SO2},
  ${payload.NO2},
  ${payload.CO},
  '${payload.id}',
  ${payload.area},
  ${payload.ts}
)
```
Then click "Test Connectivity" to make sure config is correct and click "Create" button to create the data bridge  
#### 2.3 Create Rule To Integration to TDengine
Choose "Rules" -> Click "Create":
![EMQX create_rule](docs/images/EMQX_create_rule.png "EMQX create rule")
in next page, enter below SQL:
```SQL
SELECT
  payload
FROM
  "sensor/data"
```
then click "Add Action" -> choose "Forwarding with Data Bridge" -> choose "webhook", to the data bridge you created in last step
![EMQX add_action](docs/images/EMQX_add_action.png "EMQX add action")

### 3. use mock simulator to create mock data to EMQX broker
pull this project to your local:
```shell
git clone https://github.com/freezonex/TDengine
```
enter mock script folder and run mock.js
```shell
cd docs/examples/other
node mock.js
hongzhi@hongzhis-MacBook-Air other % node mock.js    
client mock_client_0 connected
client mock_client_1 connected
client mock_client_2 connected
client mock_client_3 connected
client mock_client_4 connected
client mock_client_5 connected
client mock_client_6 connected
client mock_client_7 connected
client mock_client_8 connected
client mock_client_9 connected
10:26:33 PM send success.
```
then check EMQX broker rule dashboard and TDengine database, to check if dataflow work successfully.

### 4. Grafana Dashboard
In this step we are going to create a demo dashboard to show the data from TDengine  
#### 4.1 Setup Graphana
Install Graphana on your local system. To sign in to Graphana for the first time:
1. Go to http://localhost:3000/ (Change the actual port number base on the docker configuration)
2. On the sign in page enter the following  
*Username: admin*    
*Password: admin*  
If successful, you can see a prompt and change your password on that prompt.
#### 4.2 Integrate Graphana with TDengine
Install TDengine plugin using either **GUI** in the Graphana plugin catalog or using **script**.
* First Option - execute the following script: 
```
bash -c "$(curl -fsSL \
  https://raw.githubusercontent.com/taosdata/grafanaplugin/master/install.sh)" -- \
  -a http://localhost:6041 \
  -u root \
  -p taosdata
```
* Second Option - use GUI to install the TDengine plugin:  
![grafana_tdengine_plugin](docs/images/grafana_tdengine_plugin.png "Install TDengine plugin")
After installing the plugin, add a new datasource of TDengine inside the `connections` pannel. Configure the Host, User, and password accordingly to make it work. (see example below)
![grafana_create_new_data_source](docs/images/grafana_create_datasource.png "Create new TDengine datasource")
#### 4.3 Create Dashboard
Now you are able to create dashboards using the TDengine datasource.   
For example, import dashboard script as shown in the figure:
![grafana_dashboard_import](docs/images/grafana-dashboard-import.png "Import grafana dashboard")
choose script/dashboard_test.json file to import:
![grafana_dashboard_test](docs/images/grafana-dashboard-test.png "grafana dashboard")
also can set to dynamic refresh for demo purpose:
![grafana_dynamic](docs/images/grafana-dynamic.png "grafana dynamic")


## Deploy on server:
server ip: 192.168.31.121  
ssh to server: 
```shell
# ssh root@192.168.31.121
password: supos
```
Access noddred and EMQX from default GUI:  192.168.31.121
Access TDengine:
```shell
# docker-compose exec tdengine /bin/bash
# taos
```