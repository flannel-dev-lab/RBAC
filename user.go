package RBAC

import (
    "errors"
)

// A User represents a human being. A User can be extended to represent
// machines, networks, etc if necessary
type User struct {
    UserId  int     // should come from the underlying system
}


// (RC-04) Core RBAC: Creates a new RBAC user
func AddUser(name string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC
func DeleteUser(name string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func AssignedRoles(user User) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}
