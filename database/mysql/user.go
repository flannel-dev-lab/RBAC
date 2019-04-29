package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql"
)

// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
func (databaseService *DatabaseService) AddUser(name string) (user vars.User, err error) {

	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_user` SET `name`= ?")
	if err != nil {
		return user, err
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return user, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = int(insertId)
	user.Name = name

	return user, nil
}

// (RC-26) Core RBAC: Deletes an existing user from RBAC, Deletes Sessions and User assignments
func (databaseService *DatabaseService) DeleteUser(userId int) (bool, error) {
	// TODO Delete User Assignments and Delete Sessions
	stmt, err := databaseService.Conn.Prepare("DELETE FROM `rbac_user` WHERE `rbac_user_id`= ?")

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(userId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-09) Core RBAC: Returns a set of roles assigned to a given user
func (databaseService *DatabaseService) AssignedRoles(userId int) ([]vars.Role, error) {
	stmt, prepErr := databaseService.Conn.Prepare("SELECT `rbac_role_id` FROM `rbac_user_role` WHERE `rbac_user_id` = ?")
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
		err = result.Scan(&role.Id)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// This function returns the set of operations a given user is permitted to perform on a given
// object, obtained either directly or through his/her assigned roles.
func (databaseService *DatabaseService) UserOperationOnObject(userId, objectId int) ([]vars.Operation, error) {
	stmt, err := databaseService.Conn.Prepare("select rp.rbac_operation_id from rbac_permission rp inner join rbac_role_permission rrp on rp.rbac_permission_id=rrp.rbac_permission_id inner join rbac_user_role rur on rrp.rbac_role_id=rur.rbac_role_id where rur.rbac_user_id=? and rp.rbac_object_id=?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Query(userId, objectId)
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

