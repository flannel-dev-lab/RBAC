package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql"
)

// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
func (databaseService *DatabaseService) CreateSession(userId int, name string) (session vars.Session, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_session` SET `name`= ?, `rbac_user_id` = ?")
	if err != nil {
		return session, err
	}

	result, err := stmt.Exec(name, userId)
	if err != nil {
		return session, err
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return session, insertIdErr
	}

	session.Id = int(insertId)
	session.Name = name
	session.UserId = userId

	return session, nil
}

// (RC-23) Core RBAC: Delete a given session with a given owner user
func (databaseService *DatabaseService) DeleteSession(userId int, sessionName string) (bool, error) {
	stmt, stmtErr := databaseService.Conn.Prepare("DELETE FROM `rbac_session` WHERE `rbac_user_id`= ? AND `name` = ?")
	if stmtErr != nil {
		return false, stmtErr
	}

	_, err := stmt.Exec(userId, sessionName)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-01) Core RBAC: Add a role as an active role of a session whose owner is a given user
func (databaseService *DatabaseService) AddActiveRole(userId, sessionId, roleId int) (bool, error) {
	result := databaseService.Conn.QueryRow("select rur.rbac_user_id, rur.rbac_role_id from rbac_user_role rur inner join rbac_session rs where rs.rbac_session_id=? and rur.rbac_role_id=? and rur.rbac_user_id = ?", sessionId, roleId, userId)

	var user vars.User
	var role vars.Role

	err := result.Scan(&user.Id, &role.Id)

	if err != nil {
		return false, err
	}

	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_session_role` SET `rbac_role_id`= ?, `rbac_user_id` = ?, `rbac_session_id` = ?")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(roleId, userId, sessionId)
	if err != nil {
		return false, err
	}

	return true, err
}

// (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
func (databaseService *DatabaseService) DropActiveRole(userId, sessionId, roleId int) (bool, error) {
	stmt, stmtErr := databaseService.Conn.Prepare("DELETE FROM `rbac_session_role` WHERE `rbac_user_id`= ? AND `rbac_role_id`= ? AND `rbac_session_id` = ?")
	if stmtErr != nil {
		return false, stmtErr
	}

	_, err := stmt.Exec(userId, roleId, sessionId)
	if err != nil {
		return false, err
	}

	return true, nil
}

// (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
// or not to perform a given operation on a given object
func (databaseService *DatabaseService) CheckAccess(sessionId, operationId, objectId int) (bool, error) {
	stmt, err := databaseService.Conn.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_session rs JOIN rbac_user_role rur ON rs.rbac_user_id = rur.rbac_user_id JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rs.name = ? AND rp.rbac_object_id = ? AND rp.rbac_operation_id = ?")
	if err != nil {
		return false, err
	}

	result, err := stmt.Query(sessionId, objectId, operationId)
	if err != nil {
		return false, err
	}

	var permissions []vars.Permission

	for result.Next() {
		var permission vars.Permission
		err = result.Scan(&permission.Id, &permission.ObjectId, &permission.OperationId)
		if err != nil {
			return false, err
		}
		permissions = append(permissions, permission)
	}

	if len(permissions) > 0 {
		return true, nil
	}

	return false, nil
}