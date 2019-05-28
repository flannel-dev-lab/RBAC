package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql" // Importing mysql Driver
)

// CreateObject Create an Object
func (databaseService *DatabaseService) CreateObject(objectName, description string) (object vars.Object, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_object` SET `name`= ?, description = ?")
	if err != nil {
		return object, err
	}

	result, err := stmt.Exec(objectName, description)
	if err != nil {
		return object, err
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return object, insertIdErr
	}

	object.Id = int(insertId)
	object.Name = objectName
	object.Description = description

	return object, nil
}

// RemoveObject Removes an Object
func (databaseService *DatabaseService) RemoveObject(objectName string) (bool, error) {
	stmt, prepErr := databaseService.Conn.Prepare("DELETE FROM `rbac_object` WHERE `name` = ?")
	if prepErr != nil {
		return false, prepErr
	}

	_, err := stmt.Exec(objectName)
	if err != nil {
		return false, err
	}

	return true, nil
}
