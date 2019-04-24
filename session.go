package RBAC

import (
)

// A Session represents a user as owner and an active role set
type Session struct {
    Id              int
    Name            string
    UserId          int
}

// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
func CreateSession(userId int, name string) (Session, error) { 
    var session Session
    DbInit()

    stmt, stmtErr := DBWrite.Prepare("INSERT INTO `rbac_session` SET `name`= ?, `rbac_user_id` = ?")
    if stmtErr != nil {
        return session, stmtErr
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
func DeleteSession(user User, sessionName string) (bool, error) {
    DbInit()

    stmt, stmtErr := DBWrite.Prepare("DELETE FROM `rbac_session` WHERE `rbac_user_id`= ? AND `name` = ?")
    if stmtErr != nil {
        return false, stmtErr
    }

    _, err := stmt.Exec(user.Id, sessionName)
    if err != nil {
        return false, err
    }

    return true, nil
}

// (RC-01) Core RBAC: Add a role as an active role of a session whose owner is a given user
func AddActiveRole(user User, session Session, roleId int) (bool, error) {
    // Not implemented currently
    return true, nil
}

// (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
func DropActiveRole(user User, session Session, roleName string) (bool, error) {
    // Not implemented currently
    return true, nil
}

// (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
// or not to perform a given operation on a given object
func CheckAccess(session Session, operation Operation, object Object) (bool, error) {
    DbInit()

    stmt, prepErr := DBRead.Prepare("SELECT rp.rbac_permission_id, rp.rbac_object_id, rp.rbac_operation_id FROM rbac_session rs JOIN rbac_user_role rur ON rs.rbac_user_id = rur.rbac_user_id JOIN rbac_role_permission rrp ON rur.rbac_role_id = rrp.rbac_role_id JOIN rbac_permission rp ON rrp.rbac_permission_id = rp.rbac_permission_id WHERE rs.name = ? AND rp.rbac_object_id = ? AND rp.rbac_operation_id = ?")
    if prepErr != nil {
        return false, prepErr
    }

    result, err := stmt.Query(session.Name, object.Id, operation.Id)
    if err != nil {
        return false, err
    }

    prms := []Permission{}
    for result.Next() {
        var perm Permission
        err = result.Scan(&perm.Id, &perm.ObjectId, &perm.OperationId)
        if err != nil {
            return false, err
        }
        prms = append(prms, perm)
    }

    if len(prms) > 0 {
        return true, nil
    }
    
    return false, nil
}
