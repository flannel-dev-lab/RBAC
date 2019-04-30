//  For a system that implements RBAC, the objects ANSI INCITS 359-2004
// can represent information containers (e.g., files, directories, in an operating system,
// and/or columns, rows, tables, and views within a database management system)
package rbac

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

type RBACObject struct {
	DBService database.DatabaseService
}

// Create an Object
func (rbacObject *RBACObject) CreateObject(name, description string) (vars.Object, error) {
	return rbacObject.DBService.CreateObject(name, description)
}

// Remove an Object
func (rbacObject *RBACObject) RemoveObject(objectId int) (bool, error) {
	return rbacObject.DBService.RemoveObject(objectId)
}
