package rbac

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

// UserObject Interface to expose user operations
type UserObject struct {
	DBService database.DatabaseService
}

// AddUser (RC-04) Core RBAC: Creates a new RBAC user
func (userObject *UserObject) AddUser(userName string) (vars.User, error) {
	return userObject.DBService.AddUser(userName)
}

// DeleteUser (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (userObject *UserObject) DeleteUser(userName string) (bool, error) {
	return userObject.DBService.DeleteUser(userName)
}

// AssignedRoles (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func (userObject *UserObject) AssignedRoles(userId int) ([]vars.Role, error) {
	return userObject.DBService.AssignedRoles(userId)
}

// UserOperationOnObject This function returns the set of operations a given user is permitted to perform on a given
// object, obtained either directly or through his/her assigned roles.
func (userObject *UserObject) UserOperationOnObject(userId int, objectName string) ([]vars.Operation, error) {
	return userObject.DBService.UserOperationOnObject(userId, objectName)
}
