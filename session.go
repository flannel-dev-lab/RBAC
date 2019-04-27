package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
    "github.com/flannel-dev-lab/RBAC/vars"
)

type SessionObject struct {
    DBService database.DatabaseService
}

// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
func (sessionObject *SessionObject) CreateSession(userId int, name string) (vars.Session, error) {
    return sessionObject.DBService.CreateSession(userId, name)
}

// (RC-23) Core RBAC: Delete a given session with a given owner user
func (sessionObject *SessionObject) DeleteSession(userId int, sessionName string) (bool, error) {
    return sessionObject.DBService.DeleteSession(userId, sessionName)
}

// (RC-01) Core RBAC: Add a role as an active role of a session whose owner is a given user
/*func AddActiveRole(user User, session Session, roleId int) (bool, error) {
    // Not implemented currently
    return true, nil
}*/

// (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
// TODO
/*func DropActiveRole(user User, session Session, roleName string) (bool, error) {
    // Not implemented currently
    return true, nil
}*/

// (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
// or not to perform a given operation on a given object
// TODO
/*func CheckAccess(session Session, operation Operation, object Object) (bool, error) {
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
*/