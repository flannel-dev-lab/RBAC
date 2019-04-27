package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
)

// (RC-06) Core RBAC: Creates a new role if not exists. Duplicate roles are not allowed
func (databaseService *DatabaseService) AddRole(name, description string) (role vars.Role, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_role` SET `name`= ?, description = ?")
	if err != nil {
		return role, err
	}

	result, err := stmt.Exec(name, description)
	if err != nil {
		return role, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return role, err
	}

	role.Id = int(insertId)
	role.Name = name
	role.Description = description

	return role, nil
}

// (RC-22) Core RBAC: Deletes an existing role and deletes the role session
func (databaseService *DatabaseService) DeleteRole(roleId int) (bool, error) {
	// TODO Delete role session
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_role` WHERE `rbac_role_id`= ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(roleId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-10) Core RBAC: Assigns a user to a role, will return error if the role is already assigned to the user
func (databaseService *DatabaseService) AssignUser(userId, roleId int) (bool, error) {
	stmt, stmtErr := databaseService.Conn.Prepare("INSERT INTO `rbac_user_role` SET `rbac_user_id`= ?, `rbac_role_id` = ?")
	if stmtErr != nil {
		return false, stmtErr
	}

	_, err := stmt.Exec(userId, roleId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-18) Core RBAC: Remove a user from a role and deletes session
func (databaseService *DatabaseService) DeassignUser(userId, roleId int) (bool, error) {
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_user_role` WHERE `rbac_user_id`= ? AND `rbac_role_id` = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(userId, roleId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-11) Core RBAC: Return the set of users assigned to a given role
func (databaseService *DatabaseService) AssignedUsers(roleId int) ([]vars.User, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT `rbac_user_id` FROM `rbac_user_role` WHERE `rbac_role_id` = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(roleId)
	if err != nil {
		return nil, err
	}

	var users []vars.User
	for result.Next() {
		var user vars.User
		err = result.Scan(&user.Id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// (RC-36) Core RBAC: Return the set of active roles associated with a session
func (databaseService *DatabaseService) SessionRoles(sessionId int) ([]vars.Role, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT `rbac_role_id` FROM `rbac_session_role` WHERE `rbac_session_id` = ?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(sessionId)
	if err != nil {
		return nil, err
	}

	var roles []vars.Role
	for result.Next() {
		var role vars.Role
		err = result.Scan(&role.Id)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// This function returns the set of operations a given role is permitted to perform on a given object
func (databaseService *DatabaseService) RoleOperationOnObject(roleId, objectId int) ([]vars.Operation, error) {
	stmt, err := databaseService.Conn.Prepare("select rp.rbac_operation_id from rbac_permission rp inner join rbac_role_permission rrp on rp.rbac_permission_id=rrp.rbac_permission_id where rp.rbac_object_id=? and rrp.rbac_role_id=?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(objectId, roleId)
	if err != nil {
		return nil, err
	}

	var operations []vars.Operation
	for result.Next() {
		var operation vars.Operation
		err = result.Scan(&operation.Id)
		if err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}

	return operations, nil
}