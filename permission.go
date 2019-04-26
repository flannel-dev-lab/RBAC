package RBAC
/*
import (
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

    // Find or create a corresponding permission
    prms, err := FindPermission(ob.Id, op.Id)
    if err != nil {
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

    // Attach the permission to the role
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
// Spec deviation - accepting roleId instead of roleName
func RevokePermission(ob Object, op Operation, roleId int) (bool, error) {
    // Find a corresponding permission
    prms, err := FindPermission(ob.Id, op.Id)
    if (err != nil) {
        return false, err
    }

    stmt, stmtErr := DBWrite.Prepare("DELETE FROM `rbac_role_permission` WHERE `rbac_role_id` = ? AND `rbac_permission_id` = ?")
    if stmtErr != nil {
        return false, stmtErr
    }

    _, err = stmt.Exec(roleId, prms.Id)
    if err != nil {
        return false, err
    }

    return true, nil

}

// (RC-34) Core RBAC: Return the set of permissions granted to a given role
func RolePermissions(role Role) ([]Permission, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_role_permission rrp JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rrp.rbac_role_id = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(role.Id)
    if err != nil {
        return nil, err
    }

    prms := []Permission{}
    for result.Next() {
        var perm Permission
        err = result.Scan(&perm.Id, &perm.ObjectId, &perm.OperationId)
        if err != nil {
            return nil, err
        }
        prms = append(prms, perm)
    }

    return prms, nil
}

// (RC-43) Core RBAC: Return the set of permissions granted to a given user
func UserPermissions(user User) ([]Permission, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_user_role rur JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rur.rbac_user_id = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(user.Id)
    if err != nil {
        return nil, err
    }

    prms := []Permission{}
    for result.Next() {
        var perm Permission
        err = result.Scan(&perm.Id, &perm.ObjectId, &perm.OperationId)
        if err != nil {
            return nil, err
        }
        prms = append(prms, perm)
    }

    return prms, nil
}

// (RC-35) Core RBAC: Return the set of permissions assigned to a given session
func SessionPermissions(session Session) ([]Permission, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_session rs JOIN rbac_user_role rur ON rs.rbac_user_id = rur.rbac_user_id JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rs.name = ?")
    if prepErr != nil {
        return nil, prepErr
    }

    result, err := stmt.Query(session.Name)
    if err != nil {
        return nil, err
    }

    prms := []Permission{}
    for result.Next() {
        var perm Permission
        err = result.Scan(&perm.Id, &perm.ObjectId, &perm.OperationId)
        if err != nil {
            return nil, err
        }
        prms = append(prms, perm)
    }

    return prms, nil
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

    insertId, insertIdErr := result.LastInsertId()
    if insertIdErr != nil {
        return prms, err
    }

    prms.Id = int(insertId)
    prms.ObjectId = objectId
    prms.OperationId = operationId

    return prms, nil
}
*/