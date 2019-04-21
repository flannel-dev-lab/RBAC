// Copyright 2019 Flannel Development Laboratory. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package RBAC implements the INCITS 359-2012(R2017) standard
//
// RBAC stands for Role Based Access Control. This package is meant
// to facilitate access control at any size.
package RBAC

import (
    "errors"
)

// A Role is a job function within the context of an organization
type Role struct {
    RoleId          int
    Name            string
    Description     string
}


// (RC-06) Core RBAC: Creates a new role
func AddRole(name string, description string) (bool, error) {
    DbInit()
    
    stmt, err := DBWrite.Prepare("INSERT INTO `rbac_role` SET `name`= ?, description = ?")
    if err != nil {
        return false, err
    }

    _, err = stmt.Exec(name, description)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-22) Core RBAC: Deletes an existing role
func DeleteRole(name string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-10) Core RBAC: Assigns a user to a role
func AssignUser(user User, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-18) Core RBAC: Remove a user from a role
func DeassignUser(user User, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-11) Core RBAC: Return the set of users assigned to a given role
func AssignedUsers(role Role) ([] int, error) {
    return nil, errors.New("Not yet implemented")
}

// (RC-36) Core RBAC: Return the set of active roles associated with a session
func SessionRoles(session Session) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}
