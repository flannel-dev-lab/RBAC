package database

import (
	"errors"
	"github.com/flannel-dev-lab/RBAC/database/mysql"
	"github.com/flannel-dev-lab/RBAC/vars"
)



type DatabaseService interface {
	// Creates a DB Connection
	CreateDBConnection(driver, username, password, hostname, databaseName, port string) error

	// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
	AddUser(name string) (vars.User, error)

	// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
	DeleteUser(userId int) (bool, error)

	// (RC-06) Core RBAC: Creates a new role if not exists. Duplicate roles are not allowed
	AddRole(name, description string) (vars.Role, error)

	// (RC-22) Core RBAC: Deletes an existing role and deletes the role session
	DeleteRole(roleId int) (bool, error)

	// (RC-10) Core RBAC: Assigns a user to a role, will return error if the role is already assigned to the user
	AssignUser(userId, roleId int) (bool, error)

	// (RC-18) Core RBAC: Remove a user from a role
	DeassignUser(userId, roleId int) (bool, error)

	// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
	CreateSession(userId int, name string) (vars.Session, error)

	// (RC-23) Core RBAC: Delete a given session with a given owner user
	DeleteSession(userId int, sessionName string) (bool, error)

	// (RC-11) Core RBAC: Return the set of users assigned to a given role
	AssignedUsers(roleId int) ([]vars.User, error)

	// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
	AssignedRoles(userId int) ([]vars.Role, error)

	// Create an Object
	CreateObject(name, description string) (vars.Object, error)

	// Removes an Object
	RemoveObject(objectId int) (bool, error)

	// Add an operation
	AddOperation(name, description string) (vars.Operation, error)

	// Delete Operation
	DeleteOperation(operationId int) (bool, error)

	// Creates a new Permission
	CreatePermission(objectId, operationId int) (vars.Permission, error)

	// (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
	// Grants a role the permission to perform an operation on an object
	GrantPermission(permissionId, roleId int) (bool, error)

	// Search for existing permission record
	FindPermission(objectId int, operationId int) (vars.Permission, error)

	// (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
	// Spec deviation - accepting roleId instead of roleName
	RevokePermission(permissionId, roleId int) (bool, error)

	// (RC-34) Core RBAC: Return the set of permissions granted to a given role
	RolePermissions(roleId int) ([]vars.Permission, error)

	// (RC-43) Core RBAC: Return the set of permissions granted to a given user
	UserPermissions(userId int) ([]vars.Permission, error)

	// (RC-35) Core RBAC: Return the set of permissions assigned to a given session
	SessionPermissions(sessionId int) ([]vars.Permission, error)

	// (RC-01) Core RBAC: Add a role as an active role of a session whose owner is a given user
	AddActiveRole(userId, sessionId, roleId int) (bool, error)

	// (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
	DropActiveRole(userId, sessionId, roleId int) (bool, error)

	// (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
	// or not to perform a given operation on a given object
	CheckAccess(sessionId, operationId, objectId int) (bool, error)

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
