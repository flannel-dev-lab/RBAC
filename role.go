// Copyright 2019 Flannel Development Laboratory. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package RBAC implements the INCITS 359-2012(R2017) standard
//
// RBAC stands for Role Based Access Control. This package is meant
// to facilitate access control at any size.
package RBAC

import (
)

// A Role is a job function within the context of an organization
type Role struct {
    Id              int
    Name            string
    Description     string
}


// (RC-06) Core RBAC: Creates a new role
func AddRole(name string, description string) (Role, error) {
    var role Role

    DbInit()
    
    stmt, stmtErr := DBWrite.Prepare("INSERT INTO `rbac_role` SET `name`= ?, description = ?")
    if stmtErr != nil {
        return role, stmtErr
    }

    result, err := stmt.Exec(name, description)
    if err != nil {
        return role, err
    }

    insertId, insertIdErr := result.LastInsertId()
    if insertIdErr != nil {
        return role, insertIdErr
    }

    role.Id = int(insertId)
    role.Name = name
    role.Description = description

    return role, nil
}

// (RC-22) Core RBAC: Deletes an existing role
func DeleteRole(roleId int) (bool, error) {
    DbInit()
    
    stmt, err := DBWrite.Prepare("DELETE FROM `rbac_role` WHERE `rbac_role_id`= ?")
    if err != nil {
        return false, err
    }

    _, err = stmt.Exec(roleId)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-10) Core RBAC: Assigns a user to a role
func AssignUser(user User, roleId int) (bool, error) {
    DbInit()

    stmt, stmtErr := DBWrite.Prepare("INSERT INTO `rbac_user_role` SET `rbac_user_id`= ?, `rbac_role_id` = ?")
    if stmtErr != nil {
        return false, stmtErr
    }

    _, err := stmt.Exec(user.Id, roleId)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-18) Core RBAC: Remove a user from a role
func DeassignUser(user User, roleId int) (bool, error) {
    DbInit()

    stmt, stmtErr := DBWrite.Prepare("DELETE FROM `rbac_user_role` WHERE `rbac_user_id`= ? AND `rbac_role_id` = ?")
    if stmtErr != nil {
        return false, stmtErr
    }

    _, err := stmt.Exec(user.Id, roleId)
    if err != nil {
        return false, err
    }

    return true, nil
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

