package RBAC

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

type UserObject struct {
	DBService database.DatabaseService
}

// (RC-04) Core RBAC: Creates a new RBAC user
func (userObject *UserObject) AddUser(name string) (vars.User, error) {
	return userObject.DBService.AddUser(name)
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (userObject *UserObject) DeleteUser(userId int) (bool, error) {
	return userObject.DBService.DeleteUser(userId)
}

// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func (userObject *UserObject) AssignedRoles(userId int) ([]vars.Role, error) {
	return userObject.DBService.AssignedRoles(userId)
}
