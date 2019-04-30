// Package rbac Operation is an executable image of a program, which upon invocation executes some function for the
// user. The types of operations and objects that RBAC controls are dependent on the type
// of system in which it will be implemented. For example, within a file system, operations
// might include read, write, and execute; within a database management system, operations
// might include insert, delete, append and update
package rbac

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

// OperationObject Interface to expose rbac operations
type OperationObject struct {
	DBService database.DatabaseService
}

// AddOperation Creates a new operation
func (operationObject *OperationObject) AddOperation(name, description string) (vars.Operation, error) {
	return operationObject.DBService.AddOperation(name, description)
}

// DeleteOperation Deletes a operation
func (operationObject *OperationObject) DeleteOperation(operationId int) (bool, error) {
	return operationObject.DBService.DeleteOperation(operationId)
}
