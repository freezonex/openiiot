CREATE DATABASE IF NOT EXISTS openiiot;

USE openiiot;

CREATE TABLE IF NOT EXISTS tenant (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(200),
    is_default TINYINT(1) default false,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100),
    description VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    role VARCHAR(50) NOT NULL,     -- Viewer, Editor, Admin, SuperAdmin
    auth_id VARCHAR(100),          -- auth_id is userid from 3rd party, for example person_code in SupOS
    source VARCHAR(100),           -- openiiot: created by openiiot, supos: created by supos
    token VARCHAR(2048),
    token_updatetime TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS edge (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    url VARCHAR(300) NOT NULL,    -- http://nodered-url:1883
    username VARCHAR(100),
    password VARCHAR(100),
    type VARCHAR(100) NOT NULL,   -- nodered, restful, modbus, opcua, opcda
    source VARCHAR(100),          -- user: self-defined
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS core (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    url VARCHAR(300) NOT NULL,
    username VARCHAR(100),
    password VARCHAR(100),
    type VARCHAR(100) NOT NULL,       -- mqtt | tdengine, 
    source VARCHAR(100),              -- openiiot: created by openiiot, user: self-defined
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS app (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    url VARCHAR(300) NOT NULL,
    username VARCHAR(100),
    password VARCHAR(100),
    type VARCHAR(100) NOT NULL,       -- grafana | odm | bi, 
    source VARCHAR(100),              -- openiiot: created by openiiot, user: self-defined
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS flow (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(200),
    tenant_id BIGINT NOT NULL,
    last_modified_by VARCHAR(100),    -- username who modified the flow last time
    flow_type VARCHAR(100),           -- demo mean it is a demo flow, demo flow is global and not belong to specific tenant
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS flow_edge (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    flow_id BIGINT UNSIGNED NOT NULL,
    edge_id BIGINT UNSIGNED NOT NULL,
    script MEDIUMTEXT,           -- save script here, for example, nodered flow script
    script2 MEDIUMTEXT,
    script3 MEDIUMTEXT,
    script4 MEDIUMTEXT
);

CREATE TABLE IF NOT EXISTS flow_core (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    flow_id BIGINT UNSIGNED NOT NULL,
    core_id BIGINT UNSIGNED NOT NULL,
    script MEDIUMTEXT,           -- save script here
    script2 MEDIUMTEXT,
    script3 MEDIUMTEXT
);

CREATE TABLE IF NOT EXISTS flow_app (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    flow_id BIGINT UNSIGNED NOT NULL,
    app_id BIGINT UNSIGNED NOT NULL,
    script MEDIUMTEXT,          -- save script here
    script2 MEDIUMTEXT,
    script3 MEDIUMTEXT
);

-- CREATE TABLE IF NOT EXISTS tenant (
--     id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
--     name VARCHAR(100) NOT NULL UNIQUE,
--     value VARCHAR(200),
--     update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
--     create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

CREATE TABLE IF NOT EXISTS global_config (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    value VARCHAR(200),
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- WMS related table
-- warehouse
CREATE TABLE IF NOT EXISTS wms_warehouse (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    warehouse_id VARCHAR(100) NOT NULL UNIQUE,
    type VARCHAR(100) NOT NULL,
    manager VARCHAR(100),
    department VARCHAR(100),
    email VARCHAR(200),
    project_group VARCHAR(100),
    note VARCHAR(200),
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- warehouse storage location
CREATE TABLE IF NOT EXISTS wms_storage_location (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    warehouse_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(100) NOT NULL,
    occupied BOOLEAN default false,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- location material detail
CREATE TABLE IF NOT EXISTS wms_storage_location_material (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    store_location_id BIGINT UNSIGNED NOT NULL,
    material_id BIGINT UNSIGNED NOT NULL,
    quantity INT NOT NULL
);

-- define material basic
CREATE TABLE IF NOT EXISTS wms_material (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_code VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    storage_location_id BIGINT UNSIGNED default 0,      -- 0: not in repo yet
    product_type VARCHAR (100),
    quantity INT,
    unit VARCHAR(100),
    note VARCHAR(200),
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- add RFID to material
CREATE TABLE IF NOT EXISTS wms_rfid_material (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    rfid VARCHAR(100) NOT NULL,
    material_id BIGINT UNSIGNED NOT NULL,
    quantity INT NOT NULL,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- inbound order
CREATE TABLE IF NOT EXISTS wms_inbound (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ref_id VARCHAR(100) NOT NULL UNIQUE,                -- yyyymmdd0000
    type VARCHAR (100) NOT NULL,                        -- inbound type
    source VARCHAR(100) NOT NULL,                       -- PDA or Manual
    note VARCHAR(200),
    operator VARCHAR(100) NOT NULL,
    status VARCHAR(100),
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- inbound order detail
CREATE TABLE IF NOT EXISTS wms_inbound_record (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    inbound_id BIGINT NOT NULL,
    stock_location_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    quantity INT NOT NULL
);

-- outbound order
CREATE TABLE IF NOT EXISTS wms_outbound (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ref_id VARCHAR(100) NOT NULL UNIQUE,                -- yyyymmdd0000
    type VARCHAR (100) NOT NULL,
    source VARCHAR(100) NOT NULL,                       -- PDA or Manual
    note VARCHAR(200),
    operator VARCHAR(100) NOT NULL,
    status VARCHAR(100),
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- outbound order detail
CREATE TABLE IF NOT EXISTS wms_outbound_record(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    outbound_id BIGINT NOT NULL,
    stock_location_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    quantity INT NOT NULL
);

-- stocktaking order
CREATE TABLE IF NOT EXISTS wms_stocktaking (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ref_id VARCHAR(100) NOT NULL UNIQUE,
    type VARCHAR (100) NOT NULL,
    source VARCHAR(100) NOT NULL,                       -- PDA or Manual
    note VARCHAR(200),
    operator VARCHAR(100) NOT NULL,
    status VARCHAR(100),
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- stocktaking order detail include result: 
-- discrepancy = 0 : Inventory Balance 
-- discrepancy < 0 : Inventory Shrinkage, Actual inventory quantity is less than recorded inventory
-- discrepancy > 0 : Inventory Surplus, Actual inventory quantity is greater than recorded inventory
CREATE TABLE IF NOT EXISTS wms_stocktaking_record(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    stocktaking_id BIGINT NOT NULL,
    stock_location_id BIGINT NOT NULL,
    material_id BIGINT NOT NULL,
    quantity INT,
    stock_quantity INT,
    discrepancy INT
);
