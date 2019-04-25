// Copyright 2019 Flannel Development Laboratory. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package RBAC implements the INCITS 359-2012(R2017) standard
//
// RBAC stands for Role Based Access Control. This package is meant
// to facilitate access control at any size.
package RBAC

import "github.com/flannel-dev-lab/RBAC/database"

type RoleObject struct {
    DBService database.DatabaseService
}

// (RC-06) Core RBAC: Creates a new role
func (roleObject * RoleObject) AddRole(name string, description string) (database.Role, error) {
    return roleObject.DBService.AddRole(name, description)
}

// (RC-22) Core RBAC: Deletes an existing role and deletes the role session
func (roleObject * RoleObject) DeleteRole(roleId int) (bool, error) {
    return roleObject.DBService.DeleteRole(roleId)
}

// (RC-10) Core RBAC: Assigns a user to a role
func (roleObject * RoleObject) AssignUser(userId int, roleId int) (bool, error) {
    return roleObject.DBService.AssignUser(userId, roleId)
}

// (RC-18) Core RBAC: Remove a user from a role and deletes session
func (roleObject * RoleObject) DeassignUser(userId int, roleId int) (bool, error) {
    return roleObject.DBService.DeassignUser(userId, roleId)
}

// (RC-11) Core RBAC: Return the set of users assigned to a given role
func AssignedUsers(roleId int) ([]User, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT `rbac_user_id` FROM `rbac_user_role` WHERE `rbac_role_id` = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(roleId)
    if err != nil {
        return nil, err
    }

    users := []User{}
    for result.Next() {
        var user User
        err = result.Scan(&user.Id)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

// (RC-36) Core RBAC: Return the set of active roles associated with a session
func SessionRoles(session Session) ([]Role, error) {
    // Not implemented
    roles := []Role{}
    return roles, nil
}

