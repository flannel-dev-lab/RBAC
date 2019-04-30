# RBAC
Golang implementation of INCITS 359-2012 (R2017)

## Go version
- > 1.12



## Supported Databases
- mysql

## Testing
We are using local docker container to run unit tests

## mysql
- Make sure you have docker installed
- Add the following environment variables locally
    - `export RBAC_DB_DRIVER=mysql`
    - `export RBAC_DB_USERNAME=root`
    - `export RBAC_DB_PASSWORD=toor`
    - `export RBAC_DB_HOSTNAME=localhost`
    - `export RBAC_DB_NAME=rbac`
    - `export RBAC_DB_PORT=3306`
- Run the following command to install mysql docker container `docker build -t mysql .`
- Run the previously created container `docker container run -d -p 3306:3360 mysql`
- Go to the RBAC repository and run `go test -v`

## Installation
- Database Schema - can be found in db/schema
- Database Test Values - can be found in db/data

