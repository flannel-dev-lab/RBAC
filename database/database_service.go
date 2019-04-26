package database

import (
	"errors"
	"github.com/flannel-dev-lab/RBAC/database/mysql"
)

// A User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
	Id   int    `json:"rbac_user_id"` // should come from the underlying system
	Name string `json:"name"`         // this might need to be removed for target system
}

// A Role is a job function within the context of an organization
type Role struct {
	Id          int    `json:"rbac_role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// A Session represents a user as owner and an active role set
type Session struct {
	Id     int    `json:"rbac_session_id"`
	Name   string `json:"name"`
	UserId int    `json:"rbac_user_id"`
}

// An object can be any system resource subject to access control
type Object struct {
	Id          int    `json:"rbac_object_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DatabaseService interface {
	// Creates a DB Connection
	CreateDBConnection(driver, username, password, hostname, databaseName, port string) error

	// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
	AddUser(name string) (User, error)

	// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
	DeleteUser(userId int) (bool, error)

	// (RC-06) Core RBAC: Creates a new role if not exists. Duplicate roles are not allowed
	AddRole(name, description string) (Role, error)

	// (RC-22) Core RBAC: Deletes an existing role and deletes the role session
	DeleteRole(roleId int) (bool, error)

	// (RC-10) Core RBAC: Assigns a user to a role, will return error if the role is already assigned to the user
	AssignUser(userId, roleId int) (bool, error)

	// (RC-18) Core RBAC: Remove a user from a role
	DeassignUser(userId, roleId int) (bool, error)

	// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
	CreateSession(userId int, name string) (Session, error)

	// (RC-23) Core RBAC: Delete a given session with a given owner user
	DeleteSession(userId int, sessionName string) (bool, error)

	// (RC-11) Core RBAC: Return the set of users assigned to a given role
	AssignedUsers(roleId int) ([]User, error)

	// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
	AssignedRoles(userId int) ([]Role, error)

	// Create an Object
	CreateObject(name, description string) (Object, error)

	// Removes an Object
	RemoveObject(objectId int) (bool, error)

	// Closes a DB Connection
	CloseConnection() error
}

// Creates a Database object given the driver type
func CreateDatabaseObject(driver string) (DatabaseService, error) {
	switch driver {
	case "mysql":
		return new(mysql.DatabaseService), nil

	default:
		return nil, errors.New("unsupported driver")

	}
}
