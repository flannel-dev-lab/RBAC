package mysql

import (
	"github.com/flannel-dev-lab/RBAC/vars"
	_ "github.com/go-sql-driver/mysql"
)

// Create an Object
func (databaseService *DatabaseService) CreateObject(name, description string) (object vars.Object, err error) {
	stmt, err := databaseService.Conn.Prepare("INSERT INTO `rbac_object` SET `name`= ?, description = ?")
	if err != nil {
		return object, err
	}

	result, err := stmt.Exec(name, description)
	if err != nil {
		return object, err
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return object, insertIdErr
	}

	object.Id = int(insertId)
	object.Name = name
	object.Description = description

	return object, nil
}

// Removes an Object
func (databaseService *DatabaseService) RemoveObject(objectId int) (bool, error) {
	stmt, prepErr := databaseService.Conn.Prepare("DELETE FROM `rbac_object` WHERE `rbac_object_id` = ?")
	if prepErr != nil {
		return false, prepErr
	}

	_, err := stmt.Exec(objectId)
	if err != nil {
		return false, err
	}

	return true, nil
}
