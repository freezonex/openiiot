## quick start
bring up a local TDengine DB, build and start the service
```shell
cd tools
docker-compose up -d 
```
Create Database and Table
In this step we create the appropriate database and table schema in TDengine for receiving MQTT data. Open TDengine CLI and execute SQL bellow:
```
CREATE DATABASE test;
USE test;
CREATE TABLE sensor_data (ts TIMESTAMP, temperature FLOAT, humidity FLOAT, volume FLOAT, pm10 FLOAT, pm25 FLOAT, so2 FLOAT, no2 FLOAT, co FLOAT, sensor_id NCHAR(255), area TINYINT, coll_time TIMESTAMP);
```
Login EMQX Dashboard
Use your browser to open the URL http://IP:18083 and log in to EMQX Dashboard. The initial installation username is admin and the password is: public.
![alt text](./docs/EMQX_login.png "Title")
Edit "action"
Edit the resource configuration to add the key/value pairing for Authorization. If you use the default TDengine username and password then the value of key Authorization is:
```
Basic cm9vdDp0YW9zZGF0YQ==
```
Please refer to the TDengine REST API documentation for the authorization in details.

Enter the rule engine replacement template in the message body:
```
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
<<<<<<< HEAD
then check EMQX broker rule dashboard and TDengine database, to check if dataflow work successfully.

## if develop in mainland China, need add use domestic mirror:
```shell
$ go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
$ go env | grep GOPROXY
```
=======
## Setup Graphana
Install Graphana on your local system. To sign in to Graphana for the first time:
1. Go to http://localhost:3000/ (Change the actual port number base on the docker configuration)
2. On the sign in page enter the following  
*Username: admin*    
*Password: admin*  
If successful, you can see a prompt and change your password on that prompt.
## Integrate Graphana with TDengine
Install TDengine plugin using either **GUI** in the Graphana plugin catalog or using **script**:
```
bash -c "$(curl -fsSL \
  https://raw.githubusercontent.com/taosdata/grafanaplugin/master/install.sh)" -- \
  -a http://localhost:6041 \
  -u root \
  -p taosdata
```
If you are using GUI, after install the TDengine plugin:  
Add a new datasource of TDengine inside the `connections` pannel. Configure the Host, User, and password accordingly to make it work.
>>>>>>> 3fcfb9b405c1b88ace4990da2bde4c769d4dfb62
