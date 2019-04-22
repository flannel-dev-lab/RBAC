package RBAC

import (
)

// An object can be any system resource subject to access control
type Object struct {
    Id              int
    Name            string
    Description     string
}


// (RC-33) Core RBAC: Returns the set of operations a given role
// is permitted to perform on a given object
func RoleOperationsOnObject(role Role, object Object) ([]Operation, error) {
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
}

// Create an Object
func CreateObject(name string, description string) (Object, error) {
    var object Object

    DbInit()
    
    stmt, prepErr := DBWrite.Prepare("INSERT INTO `rbac_object` SET `name`= ?, description = ?")
    if prepErr != nil {
        return object, prepErr
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

// Remove an Object
func RemoveObject(object Object) (bool, error) {
    DbInit()

    stmt, prepErr := DBWrite.Prepare("DELETE FROM `rbac_object` WHERE `rbac_object_id` = ?")
    if prepErr != nil {
        return false, prepErr
    }

    _, err := stmt.Exec(object.Id)
    if err != nil {
        return false, err
    }

    return true, nil
}
