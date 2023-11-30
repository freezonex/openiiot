# Database

### Overview

##### Environments

PRODUCTION [SG]
- [Autoduck](https://cloud-i18n.bytedance.net/rds/detail/db/alisg/autoduck_db/overview)

BOEi18n
- [Autoduck](https://cloud-boe.bytedance.net/rds/detail/db/boei18n/autoduck_db/overview)

LOCAL
- MySQL with Docker
    - GORM automatic code generation 
    - Testing/Debugging in an offline controlled environment
- SQLITE
    - CI testing


### Local Setup

There are several use cases for local setup:
- Automatically generate GORM models in the source code
- Perform sanity checks and verification of database design and logic without affecting online data
- Setup offline controlled database state to debug and reproduce bugs safely.

##### Automatically generate GORM models in the source code

First we will need to start our local db using docker by running the following commands
```shell

# enter tools dir from repository root dir
cd tools

# run docker images and detach process(es) to background
docker-compose up -d

# if need to sync the tables from boei18n autoduck_db, run below commands
sh dump_autoduck_db_boei18n_to_docker_db.sh

```

You can access your local database via PHPMyAdmin tool on `localhost:8080` with the following credentials:
```
username: root
password: root1234
```

You can then create your tables locally for new or existing databases.

Then we have to configure some files:
1. `biz/config/conf/<env>.yml`
2. `biz/config/config.go`
3. `biz/dal/mysql/mysql.go`
4. `cmd/generate/generate.go`


If you have any configs such as DSN or PSM for a new database, add it to the corresponding env file such as boei18n `biz/config/conf/boei18n.yml` under `DBConfig`.

Then update `biz/config/config.go` to parse/unmarshal the config yml files if necessary.

Update and setup database with gorm and `MySQL` struct in `Init` function.

Update the `main` function in `cmd/generate/generate.go` to generate the GORM models automatically.

Enter the following command to run code generation on GORM models.
```shell
sh generate.sh
```

##### Perform sanity checks and verification of database design and logic without affecting online data

This is simply configuring the local database how you want to, particularly in implementing/integrating features.

##### Setup offline controlled database stateto debug and reproduce bugs safely

Same as above, this is for configuring the local database to reproduce/analyze bugs.


### CI Setup

This is needed to write test cases that interface with any database. We use SQLITE as an embedded database for safe testing decoupled from deployed databases.
Update `biz/dal/mysql/mysql.go` to connect database to sqlite based on config data.

Example:

```go

func Init(c *config.DBConfig) (*MySQL, error) {
	var m MySQL
	// CI environment
	if c.Ugqa.CI || c.Autoduck.CI {
		if c.Ugqa.CI {
			dbUgqa = connectSQLiteUgqa().Debug()
			m.DBUgqaQuery = query_ugqa.Use(dbUgqa)
		}
		if c.Autoduck.CI {
			dbAutoduck = connectSQLiteAutoduck().Debug()
			m.DBAutoduckQuery = query_autoduck.Use(dbAutoduck)
		}
		return &m, nil
	}
    ...

```
