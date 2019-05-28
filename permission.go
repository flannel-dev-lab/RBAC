package rbac

import (
	"database/sql"
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

// PermissionObject Exposes permission methods
type PermissionObject struct {
	DBService database.DatabaseService
}

// CreatePermission Create a new permission record
func (permissionObject *PermissionObject) CreatePermission(objectName, operationName string) (vars.Permission, error) {
	return permissionObject.DBService.CreatePermission(objectName, operationName)
}

// FindPermission Search for existing permission record
func (permissionObject *PermissionObject) FindPermission(objectName, operationName string) (vars.Permission, error) {
	return permissionObject.DBService.FindPermission(objectName, operationName)
}

// GrantPermission (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
// Grants a role the permission to perform an operation on an object
func (permissionObject *PermissionObject) GrantPermission(objectName, operationName, roleName string) (bool, error) {

	var permission vars.Permission
	// Find or create a corresponding permission
	permission, err := permissionObject.FindPermission(objectName, operationName)

	if err != nil {
		if err == sql.ErrNoRows {
			// Create a new permission record if one couldn't be found
			permission, err = permissionObject.CreatePermission(objectName, operationName)
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}

	return permissionObject.DBService.GrantPermission(permission.Id, roleName)
}

// RevokePermission (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
// Spec deviation - accepting roleId instead of roleName
func (permissionObject *PermissionObject) RevokePermission(objectName, operationName, roleName string) (bool, error) {
	// Find a corresponding permission
	permission, err := permissionObject.FindPermission(objectName, operationName)
	if err != nil {
		return false, err
	}
	return permissionObject.DBService.RevokePermission(permission.Id, roleName)
}

// RolePermissions (RC-34) Core RBAC: Return the set of permissions granted to a given role
func (permissionObject *PermissionObject) RolePermissions(roleName string) ([]vars.Permission, error) {
	return permissionObject.DBService.RolePermissions(roleName)
}

// UserPermissions (RC-43) Core RBAC: Return the set of permissions granted to a given user
func (permissionObject *PermissionObject) UserPermissions(userId int) ([]vars.Permission, error) {
	return permissionObject.DBService.UserPermissions(userId)
}

// SessionPermissions (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func (permissionObject *PermissionObject) SessionPermissions(sessionName string) ([]vars.Permission, error) {
	return permissionObject.DBService.SessionPermissions(sessionName)
}
