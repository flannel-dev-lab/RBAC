//  For a system that implements RBAC, the objects ANSI INCITS 359-2004
// can represent information containers (e.g., files, directories, in an operating system,
// and/or columns, rows, tables, and views within a database management system)
package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
    "github.com/flannel-dev-lab/RBAC/vars"
)

type RBACObject struct {
    DBService database.DatabaseService
}

// (RC-33) Core RBAC: Returns the set of operations a given role
// is permitted to perform on a given object
/*func RoleOperationsOnObject(role Role, object Object) ([]Operation, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT ro.rbac_operation_id, ro.name, ro.`description` FROM rbac_role_permission rrp JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id JOIN rbac_operation ro ON rp.rbac_operation_id = ro.rbac_operation_id WHERE rrp.rbac_role_id = ? AND rp.rbac_object_id = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(role.Id, object.Id)
    if err != nil {
        return nil, err
    }

    ops := []Operation{}
    for result.Next() {
        var op Operation
        err = result.Scan(&op.Id, &op.Name, &op.Description)
        if err != nil {
            return nil, err
        }
        ops = append(ops, op)
    }

    return ops, nil
}

// (RC-42) Core RBAC: Returns the set of operations a given user
// is permitted to perform on a given object
func UserOperationsOnObject(user User, object Object) ([]Operation, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT ro.rbac_operation_id, ro.name, ro.`description` FROM rbac_user_role rur JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id JOIN rbac_operation ro ON rp.rbac_operation_id = ro.rbac_operation_id WHERE rur.rbac_user_id = ? AND rp.rbac_object_id = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(user.Id, object.Id)
    if err != nil {
        return nil, err
    }

    ops := []Operation{}
    for result.Next() {
        var op Operation
        err = result.Scan(&op.Id, &op.Name, &op.Description)
        if err != nil {
            return nil, err
        }
        ops = append(ops, op)
    }

    return ops, nil
}*/

// Create an Object
func (rbacObject *RBACObject) CreateObject(name, description string) (vars.Object, error) {
    return rbacObject.DBService.CreateObject(name, description)
}

// Remove an Object
func (rbacObject *RBACObject) RemoveObject(objectId int) (bool, error) {
    return rbacObject.DBService.RemoveObject(objectId)
}
