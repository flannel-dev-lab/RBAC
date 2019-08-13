package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql" // Importing mysql Driver
)

// AddUser (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
func (databaseService *DatabaseService) AddUser(userName string) (user vars.User, err error) {

	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_user` SET `name`= ?")
	if err != nil {
		return user, err
	}

	result, err := stmt.Exec(userName)
	if err != nil {
		return user, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = int(insertId)
	user.Name = userName

	return user, nil
}

// DeleteUser (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (databaseService *DatabaseService) DeleteUser(userName string) (bool, error) {
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_user` WHERE `name`= ?")

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(userName)
	if err != nil {
		return false, err
	}

	return true, nil
}

// AssignedRoles (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func (databaseService *DatabaseService) AssignedRoles(userId int) ([]vars.Role, error) {
	stmt, prepErr := databaseService.Conn.Prepare("SELECT `rbac_role_name` FROM `rbac_user_role` WHERE `rbac_user_id` = ?")
	if prepErr != nil {
		return nil, prepErr
	}

	result, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}

	var roles []vars.Role
	for result.Next() {
		var role vars.Role
		err = result.Scan(&role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// UserOperationOnObject This function returns the set of operations a given user is permitted to perform on a given
// object, obtained either directly or through his/her assigned roles.
func (databaseService *DatabaseService) UserOperationOnObject(userId int, objectName string) ([]vars.Operation, error) {
	stmt, err := databaseService.Conn.Prepare("select distinct rp.rbac_operation_name from rbac_permission rp inner join rbac_role_permission rrp on rp.rbac_permission_id=rrp.rbac_permission_id inner join rbac_user_role rur on rrp.rbac_role_name=rur.rbac_role_name where rur.rbac_user_id=? and rp.rbac_object_name=?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(userId, objectName)
	if err != nil {
		return nil, err
	}

	var operations []vars.Operation
	for result.Next() {
		var operation vars.Operation
		err = result.Scan(&operation.Name)
		if err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}

	return operations, nil
}
