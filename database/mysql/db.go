package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseService struct {
	Conn *sql.DB
}

// Creates a DB Connection with the Database
func (databaseService *DatabaseService) CreateDBConnection(driver, username, password, hostname, databaseName, port string) error {
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

// Closes the DB Connection
func (databaseService *DatabaseService) CloseConnection() error {
	return databaseService.Conn.Close()
}
