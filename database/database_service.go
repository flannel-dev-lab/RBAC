package database

import (
    "RBAC/database/mysql"
    "errors"
)


type DatabaseService interface {
    // Creates a DB Connection
    CreateDBConnection(driver, username, password, hostname, databaseName, port string) error

    // Closes a DB Connection
    CloseConnection() error
}


// Creates a Database object given the driver type
func CreateDatabaseObject(driver string) (DatabaseService, error) {
    switch driver {
        case "mysql":
            return new(mysql.DBService), nil

        default:
            return nil, errors.New("Unsupported Driver")

    }
}
