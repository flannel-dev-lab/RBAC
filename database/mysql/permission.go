package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql" // Importing mysql Driver
)

// CreatePermission Creates a new Permission
func (databaseService *DatabaseService) CreatePermission(objectName, operationName string) (permission vars.Permission, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_permission` SET `rbac_object_name` = ?, `rbac_operation_name` = ?")
	if err != nil {
		return permission, err
	}

	result, err := stmt.Exec(objectName, operationName)
	if err != nil {
		return permission, err
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return permission, err
	}

	permission.Id = int(insertId)
	permission.ObjectName = objectName
	permission.OperationName = operationName

	return permission, nil
}

// GrantPermission (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
// Grants a role the permission to perform an operation on an object
func (databaseService *DatabaseService) GrantPermission(permissionId int, roleName string) (bool, error) {
	// Attach the permission to the role
	stmt, stmtErr := databaseService.Conn.Prepare("INSERT INTO `rbac_role_permission` SET `rbac_role_name` = ?, `rbac_permission_id` = ?")
	if stmtErr != nil {
		return false, stmtErr
	}

	_, err := stmt.Exec(roleName, permissionId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// FindPermission Search for existing permission record
func (databaseService *DatabaseService) FindPermission(objectName, operationName string) (permission vars.Permission, err error) {
	result := databaseService.Conn.QueryRow("SELECT `rbac_permission_id`, `rbac_object_name`, `rbac_operation_name` FROM `rbac_permission` WHERE `rbac_object_name` = ? AND `rbac_operation_name` = ?", objectName, operationName)

	err = result.Scan(&permission.Id, &permission.ObjectName, &permission.OperationName)

	if err != nil {
		return permission, err
	}

	return permission, nil
}

// RevokePermission (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
// Spec deviation - accepting roleId instead of roleName
func (databaseService *DatabaseService) RevokePermission(permissionId int, roleName string) (bool, error) {
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_role_permission` WHERE `rbac_role_name` = ? AND `rbac_permission_id` = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(roleName, permissionId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// RolePermissions (RC-34) Core RBAC: Return the set of permissions granted to a given role
func (databaseService *DatabaseService) RolePermissions(roleName string) ([]vars.Permission, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_name, rp.rbac_operation_name FROM rbac_role_permission rrp JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rrp.rbac_role_name = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(roleName)
	if err != nil {
		return nil, err
	}

	var permissions []vars.Permission
	for result.Next() {
		var permission vars.Permission
		err = result.Scan(&permission.Id, &permission.ObjectName, &permission.OperationName)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil

}

// UserPermissions (RC-43) Core RBAC: Return the set of permissions granted to a given user
func (databaseService *DatabaseService) UserPermissions(userId int) ([]vars.Permission, error) {
	stmt, prepErr := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_name, rp.rbac_operation_name FROM rbac_user_role rur JOIN rbac_role_permission rrp ON rur.rbac_role_name = rrp.rbac_role_name JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rur.rbac_user_id = ?")
	if prepErr != nil {
		return nil, prepErr
	}

	result, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	var permissions []vars.Permission
	for result.Next() {
		var permission vars.Permission
		err = result.Scan(&permission.Id, &permission.ObjectName, &permission.OperationName)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

// SessionPermissions (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func (databaseService *DatabaseService) SessionPermissions(sessionName string) ([]vars.Permission, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_name, rp.rbac_operation_name FROM rbac_session rs JOIN rbac_user_role rur ON rs.rbac_user_id = rur.rbac_user_id JOIN rbac_role_permission rrp ON rur.rbac_role_name = rrp.rbac_role_name JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rs.name = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(sessionName)
	if err != nil {
		return nil, err
	}

	var permissions []vars.Permission
	for result.Next() {
		var permission vars.Permission
		err = result.Scan(&permission.Id, &permission.ObjectName, &permission.OperationName)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
