package mysql

import (
	"database/sql"
	"fmt"
	"github.com/flannel-dev-lab/RBAC/database"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseService struct {
	Conn *sql.DB
}

// Creates a DB Connection with the Database
func (databaseService *DatabaseService) CreateDBConnection(driver, username, password, hostname, databaseName, port string) error {
	dbConnection, err := sql.Open(
		driver,
		fmt.Sprintf("%s:%s@tcp(%s)/%s",
			username,
			password,
			hostname,
			databaseName,
		))

	if err != nil {
		return err
	}
	databaseService.Conn = dbConnection
	return nil
}

// (RC-04) Core RBAC: Creates a new RBAC user.  The User will not carry any sessions during the creation
func (databaseService *DatabaseService) AddUser(name string) (user database.User, err error) {

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

// (RC-06) Core RBAC: Creates a new role if not exists. Duplicate roles are not allowed
func (databaseService *DatabaseService) AddRole(name string, description string) (role database.Role, err error) {
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
func (databaseService *DatabaseService) AssignUser(userId int, roleId int) (bool, error) {
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
func (databaseService *DatabaseService) DeassignUser(userId int, roleId int) (bool, error) {
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

// Closes the DB Connection
func (databaseService *DatabaseService) CloseConnection() error {
	return databaseService.Conn.Close()
}
