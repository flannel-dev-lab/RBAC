package RBAC

import (
    "errors"
    "database/sql"
)

// A Permission is a a combination of object & operation that can be enforced
type Permission struct {
    Id              int
    ObjectId        int
    OperationId     int
}


// (RC-31) Core RBAC: Grant a role a permission - must pair an object and an operation
// Spec deviation - accepting roleId instead of roleName
func GrantPermission(ob Object, op Operation, roleId int) (bool, error) {
    DbInit()

    // Find a corresponding permission
    prms, err := FindPermission(ob.Id, op.Id)
    if (err != nil) {
        if err == sql.ErrNoRows {
            // Create a new permission record if one couldn't be found
            prms, err = CreatePermission(ob.Id, op.Id)
            if (err != nil) {
                return false, err
            }
        } else {
            return false, err
        }
    }

    stmt, stmtErr := DBWrite.Prepare("INSERT INTO `rbac_role_permission` SET `rbac_role_id` = ?, `rbac_permission_id` = ?")
    if stmtErr != nil {
        return false, stmtErr
    }

    _, err = stmt.Exec(roleId, prms.Id)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-32) Core RBAC: Revoke a permission from a role - must pair an object and an operation
func RevokePermission(object Object, operation Operation, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-34) Core RBAC: Return the set of permissions granted to a given role
func RolePermissions(role Role) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// (RC-43) Core RBAC: Return the set of permissions granted to a given user
func UserPermissions(user User) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func SessionPermissions(session Session) ([]int, error) {
    return nil, errors.New("Not yet implemented")
}

// Search for existing permission record
func FindPermission(objectId int, operationId int) (Permission, error) {
    DbInit()

    result := DBRead.QueryRow("SELECT `rbac_permission_id`, `rbac_object_id`, `rbac_operation_id` FROM `rbac_permission` WHERE `rbac_object_id` = ? AND `rbac_operation_id` = ?", objectId, operationId)

    var prms Permission
    err := result.Scan(&prms.Id, &prms.ObjectId, &prms.OperationId)

    if err != nil {
        return prms, err
    }

    return prms, nil
}

// Create a new permission record
func CreatePermission(objectId int, operationId int) (Permission, error) {
    var prms Permission

    DbInit()

    stmt, stmtErr := DBWrite.Prepare("INSERT INTO `rbac_permission` SET `rbac_object_id` = ?, `rbac_operation_id` = ?")
    if stmtErr != nil {
        return prms, stmtErr
    }
    
    result, err := stmt.Exec(objectId, operationId)
    if err != nil {
        return prms, err
    }

    newId, insertIdErr := result.LastInsertId()
    if insertIdErr != nil {
        return prms, err
    }

    prms.Id = int(newId)
    prms.ObjectId = objectId
    prms.OperationId = operationId

    return prms, nil
}
