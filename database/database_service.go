package database

import (
	"errors"
	"github.com/flannel-dev-lab/RBAC/database/mysql"
)

// A User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
	Id   int64  // should come from the underlying system
	Name string // this might need to be removed for target system
}

type DatabaseService interface {
	// Creates a DB Connection
	CreateDBConnection(driver, username, password, hostname, databaseName, port string) error

	// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
	AddUser(name string) (user User, err error)

	// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
	DeleteUser(userId int64) (bool, error)

	// Closes a DB Connection
	CloseConnection() error
}

// Creates a Database object given the driver type
func CreateDatabaseObject(driver string) (DatabaseService, error) {
	switch driver {
	case "mysql":
		return new(mysql.DatabaseService), nil

	default:
		return nil, errors.New("Unsupported Driver")

	}
}
