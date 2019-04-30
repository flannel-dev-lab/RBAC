package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Importing mysql Driver
)

// DatabaseService Interface to expose DB methods
type DatabaseService struct {
	Conn *sql.DB
}

// CreateDBConnection Creates a DB Connection with the Database
func (databaseService *DatabaseService) CreateDBConnection(driver, username, password, hostname, databaseName string) error {
	dbConnection, err := sql.Open(
		driver,
		fmt.Sprintf("%s:%s@tcp(%s)/%s",
			username,
			password,
			hostname,
			databaseName,
		))

	if err != nil {
		return err
	}
	databaseService.Conn = dbConnection
	return nil
}

// CloseConnection Closes the DB Connection
func (databaseService *DatabaseService) CloseConnection() error {
	return databaseService.Conn.Close()
}
