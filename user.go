package RBAC

import (
    "errors"
)

// A User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
    Id  int     // should come from the underlying system
    Name    string  // this might need to be removed for target system
}


// (RC-04) Core RBAC: Creates a new RBAC user
func AddUser(name string) (bool, error) {
    DbInit()

    stmt, err := DBWrite.Prepare("INSERT INTO `rbac_user` SET `name`= ?")
    if err != nil {
        return false, err
    }

    _, err = stmt.Exec(name)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC
func DeleteUser(name string) (bool, error) {
    DbInit()

    stmt, err := DBWrite.Prepare("DELETE FROM `rbac_user` WHERE `name`= ?")
    if err != nil {
        return false, err
    }

    _, err = stmt.Exec(name)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func AssignedRoles(user User) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}
