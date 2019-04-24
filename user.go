package RBAC

import (
)

// A User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
    Id  int     // should come from the underlying system
    Name    string  // this might need to be removed for target system
}


// (RC-04) Core RBAC: Creates a new RBAC user
func AddUser(name string) (User, error) {
    var user User
    DbInit()

    stmt, stmtErr := DBWrite.Prepare("INSERT INTO `rbac_user` SET `name`= ?")
    if stmtErr != nil {
        return user, stmtErr
    }

    result, err := stmt.Exec(name)
    if err != nil {
        return user, err
    }

    insertId, insertIdErr := result.LastInsertId()
    if insertIdErr != nil {
        return user, insertIdErr
    }

    user.Id = int(insertId)
    user.Name = name

    return user, nil
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC
func DeleteUser(userId int) (bool, error) {
    DbInit()

    stmt, err := DBWrite.Prepare("DELETE FROM `rbac_user` WHERE `rbac_user_id`= ?")
    if err != nil {
        return false, err
    }

    _, err = stmt.Exec(userId)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func AssignedRoles(userId int) ([]Role, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT `rbac_role_id` FROM `rbac_user_role` WHERE `rbac_user_id` = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(userId)
    if err != nil {
        return nil, err
    }

    roles := []Role{}
    for result.Next() {
        var role Role
        err = result.Scan(&role.Id)
        if err != nil {
            return nil, err
        }
        roles = append(roles, role)
    }

    return roles, nil    
}
