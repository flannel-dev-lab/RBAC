package mysql

import "github.com/flannel-dev-lab/RBAC/database"

// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
func (databaseService *DatabaseService) CreateSession(userId int, name string) (session database.Session, err error) {
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
