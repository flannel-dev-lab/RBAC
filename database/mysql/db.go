package mysql

import (
	"database/sql"
	"fmt"
	"github.com/flannel-dev-lab/RBAC/database"
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

// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
func (databaseService *DatabaseService) AddUser(name string) (user database.User, err error) {

	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_user` SET `name`= ?")
	if err != nil {
		return user, err
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return user, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = insertId
	user.Name = name

	return user, nil
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (databaseService *DatabaseService) DeleteUser(userId int64) (bool, error) {
    // TODO Delete User Assignments and Delete Sessions
    stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_user` WHERE `rbac_user_id`= ?")

    if err != nil {
        return false, err
    }

    _, err = stmt.Exec(userId)
    if err != nil {
        return false, err
    }

    return true, nil
}

// Closes the DB Connection
func (databaseService *DatabaseService) CloseConnection() error {
	return databaseService.Conn.Close()
}
