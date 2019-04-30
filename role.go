// Package rbac Copyright 2019 Flannel Development Laboratory. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package RBAC implements the INCITS 359-2012(R2017) standard
//
// RBAC stands for Role Based Access Control. This package is meant
// to facilitate access control at any size.
package rbac

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

// RoleObject Exposes permission methods
type RoleObject struct {
	DBService database.DatabaseService
}

// AddRole (RC-06) Core RBAC: Creates a new role
func (roleObject *RoleObject) AddRole(name, description string) (vars.Role, error) {
	return roleObject.DBService.AddRole(name, description)
}

// DeleteRole (RC-22) Core RBAC: Deletes an existing role and deletes the role session
func (roleObject *RoleObject) DeleteRole(roleId int) (bool, error) {
	return roleObject.DBService.DeleteRole(roleId)
}

// AssignUser (RC-10) Core RBAC: Assigns a user to a role
func (roleObject *RoleObject) AssignUser(userId int, roleId int) (bool, error) {
	return roleObject.DBService.AssignUser(userId, roleId)
}

// DeassignUser (RC-18) Core RBAC: Remove a user from a role and deletes session
func (roleObject *RoleObject) DeassignUser(userId int, roleId int) (bool, error) {
	return roleObject.DBService.DeassignUser(userId, roleId)
}

// AssignedUsers (RC-11) Core RBAC: Return the set of users assigned to a given role
func (roleObject *RoleObject) AssignedUsers(roleId int) ([]vars.User, error) {
	return roleObject.DBService.AssignedUsers(roleId)
}

// SessionRoles (RC-36) Core RBAC: Return the set of active roles associated with a session
func (roleObject *RoleObject) SessionRoles(sessionId int) ([]vars.Role, error) {
	return roleObject.DBService.SessionRoles(sessionId)
}

// RoleOperationOnObject This function returns the set of operations a given role is permitted to perform on a given object
func (roleObject *RoleObject) RoleOperationOnObject(roleId, objectId int) ([]vars.Operation, error) {
	return roleObject.DBService.RoleOperationOnObject(roleId, objectId)
}
