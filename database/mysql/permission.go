package mysql

import "github.com/flannel-dev-lab/RBAC/vars"

// Creates a new Permission
func (databaseService *DatabaseService) CreatePermission(objectId, operationId int) (permission vars.Permission, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_permission` SET `rbac_object_id` = ?, `rbac_operation_id` = ?")
	if err != nil {
		return permission, err
	}

	result, err := stmt.Exec(objectId, operationId)
	if err != nil {
		return permission, err
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return permission, err
	}

	permission.Id = int(insertId)
	permission.ObjectId = objectId
	permission.OperationId = operationId

	return permission, nil
}

// (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
// Grants a role the permission to perform an operation on an object
func (databaseService *DatabaseService) GrantPermission(permissionId, roleId int) (bool, error) {
	// Attach the permission to the role
	stmt, stmtErr := databaseService.Conn.Prepare("INSERT INTO `rbac_role_permission` SET `rbac_role_id` = ?, `rbac_permission_id` = ?")
	if stmtErr != nil {
		return false, stmtErr
	}

	_, err := stmt.Exec(roleId, permissionId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Search for existing permission record
func (databaseService *DatabaseService) FindPermission(objectId int, operationId int) (permission vars.Permission, err error) {
	result := databaseService.Conn.QueryRow("SELECT `rbac_permission_id`, `rbac_object_id`, `rbac_operation_id` FROM `rbac_permission` WHERE `rbac_object_id` = ? AND `rbac_operation_id` = ?", objectId, operationId)

	err = result.Scan(&permission.Id, &permission.ObjectId, &permission.OperationId)

	if err != nil {
		return permission, err
	}

	return permission, nil
}

// (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
// Spec deviation - accepting roleId instead of roleName
func (databaseService *DatabaseService) RevokePermission(permissionId, roleId int) (bool, error) {
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_role_permission` WHERE `rbac_role_id` = ? AND `rbac_permission_id` = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(roleId, permissionId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-34) Core RBAC: Return the set of permissions granted to a given role
func (databaseService *DatabaseService) RolePermissions(roleId int) ([]vars.Permission, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_role_permission rrp JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rrp.rbac_role_id = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(roleId)
	if err != nil {
		return nil, err
	}

	var permissions []vars.Permission
	for result.Next() {
		var permission vars.Permission
		err = result.Scan(&permission.Id, &permission.ObjectId, &permission.OperationId)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil

}

// (RC-43) Core RBAC: Return the set of permissions granted to a given user
func (databaseService *DatabaseService) UserPermissions(userId int) ([]vars.Permission, error) {
	stmt, prepErr := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_user_role rur JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rur.rbac_user_id = ?")
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
		err = result.Scan(&permission.Id, &permission.ObjectId, &permission.OperationId)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}

// (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func (databaseService *DatabaseService) SessionPermissions(sessionId int) ([]vars.Permission, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_session rs JOIN rbac_user_role rur ON rs.rbac_user_id = rur.rbac_user_id JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rs.name = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(sessionId)
	if err != nil {
		return nil, err
	}

	var permissions []vars.Permission
	for result.Next() {
		var permission vars.Permission
		err = result.Scan(&permission.Id, &permission.ObjectId, &permission.OperationId)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}