package database

import (
	"errors"
	"github.com/flannel-dev-lab/RBAC/database/mysql"
)

// A User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
	Id   int  // should come from the underlying system
	Name string // this might need to be removed for target system
}

// A Role is a job function within the context of an organization
type Role struct {
	Id              int
	Name            string
	Description     string
}

type DatabaseService interface {
	// Creates a DB Connection
	CreateDBConnection(driver, username, password, hostname, databaseName, port string) error

	// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
	AddUser(name string) (User, error)

	// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
	DeleteUser(userId int) (bool, error)

	// (RC-06) Core RBAC: Creates a new role if not exists. Duplicate roles are not allowed
	AddRole(name string, description string) (Role, error)

	// (RC-22) Core RBAC: Deletes an existing role and deletes the role session
	DeleteRole(roleId int) (bool, error)

	// (RC-10) Core RBAC: Assigns a user to a role, will return error if the role is already assigned to the user
	AssignUser(userId int, roleId int) (bool, error)

	// (RC-18) Core RBAC: Remove a user from a role
	DeassignUser(userId int, roleId int) (bool, error)

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
