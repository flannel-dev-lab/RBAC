package mysql

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

type DBService struct {
    Conn *sql.DB
}

const (

)

// Creates a DB Connection with the FMD Database
func (DBService *DBService) CreateDBConnection(driver, username, password, hostname, databaseName, port string) error {
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
    DBService.Conn = dbConnection
    return nil
}

func (DBService *DBService) CloseConnection() error {
    return DBService.Conn.Close()
}
