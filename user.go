package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
)

type UserObject struct {
    DBService database.DatabaseService
    User database.User
}


// (RC-04) Core RBAC: Creates a new RBAC user
func (userObject *UserObject) AddUser(name string) (err error) {
    userInfo, err := userObject.DBService.AddUser(name)

    if err != nil {
        return err
    }

    userObject.User = userInfo
    return nil
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (userObject *UserObject) DeleteUser(userId int64) (bool, error) {
    status, err := userObject.DBService.DeleteUser(userId)

    if err != nil {
        return false, err
    }

    userObject.User = database.User{}

    return status, nil

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
