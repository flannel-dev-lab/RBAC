package RBAC

import ( 
    "errors"
)

// An object can be any system resource subject to access control
type Object struct {
    ObjectId        int
    Name            string
    Description     string
}


// (RC-33) Core RBAC: Returns the set of operations a given role
// is permitted to perform on a given object
func RoleOperationsOnObject(role Role, object Object) ([]int, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-42) Core RBAC: Returns the set of operations a given user
// is permitted to perform on a give object
func UserOperationsOnObject(user User, object Object) ([]int, error) {
    return false, errors.New("Not yet implemented")
}
