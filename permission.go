package rbac

import (
    "database/sql"
    "github.com/flannel-dev-lab/RBAC/database"
    "github.com/flannel-dev-lab/RBAC/vars"
)

type PermissionObject struct {
    DBService database.DatabaseService
}

// Create a new permission record
func (permissionObject *PermissionObject) CreatePermission(objectId, operationId int) (vars.Permission, error) {
    return permissionObject.DBService.CreatePermission(objectId, operationId)
}

// Search for existing permission record
func (permissionObject *PermissionObject) FindPermission(objectId int, operationId int) (vars.Permission, error) {
    return permissionObject.DBService.FindPermission(objectId, operationId)
}

// (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
// Grants a role the permission to perform an operation on an object
func (permissionObject *PermissionObject) GrantPermission(objectId, operationId, roleId int) (bool, error) {

    var permission vars.Permission
    // Find or create a corresponding permission
    permission, err := permissionObject.FindPermission(objectId, operationId)

    if err != nil {
        if err == sql.ErrNoRows {
            // Create a new permission record if one couldn't be found
            permission, err = permissionObject.CreatePermission(objectId, operationId)
            if err != nil {
                return false, err
            }
        } else {
            return false, err
        }
    }

    return permissionObject.DBService.GrantPermission(permission.Id, roleId)
}

// (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
// Spec deviation - accepting roleId instead of roleName
func (permissionObject *PermissionObject) RevokePermission(objectId, operationId, roleId int) (bool, error) {
    // Find a corresponding permission
    permission, err := permissionObject.FindPermission(objectId, operationId)
    if err != nil {
        return false, err
    }
    return permissionObject.DBService.RevokePermission(permission.Id, roleId)
}

// (RC-34) Core RBAC: Return the set of permissions granted to a given role
func (permissionObject *PermissionObject) RolePermissions(roleId int) ([]vars.Permission, error) {
    return permissionObject.DBService.RolePermissions(roleId)
}

// (RC-43) Core RBAC: Return the set of permissions granted to a given user
func (permissionObject *PermissionObject) UserPermissions(userId int) ([]vars.Permission, error) {
    return permissionObject.DBService.UserPermissions(userId)
}

// (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func (permissionObject *PermissionObject) SessionPermissions(sessionId int) ([]vars.Permission, error) {
    return permissionObject.DBService.SessionPermissions(sessionId)
}

