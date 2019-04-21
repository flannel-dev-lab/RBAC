package RBAC

import (
    "errors"
)

type Permission struct {
    PermissionId    int
    Name            string
    Description     string
}

// (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
func GrantPermission(object Object, operation Operation, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
func RevokePermission(object Object, operation Operation, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-34) Core RBAC: Return the set of permissions granted to a given role
func RolePermissions(role Role) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// (RC-43) Core RBAC: Return the set of permissions granted to a given user
func UserPermissions(user User) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func SessionPermissions(session Session) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}
