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
