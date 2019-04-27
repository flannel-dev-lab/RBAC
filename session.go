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
func (sessionObject *SessionObject) AddActiveRole(userId, sessionId, roleId int) (bool, error) {
    return sessionObject.DBService.AddActiveRole(userId, sessionId, roleId)
}

// (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
func (sessionObject *SessionObject) DropActiveRole(userId, sessionId, roleId int) (bool, error) {
    return sessionObject.DBService.DropActiveRole(userId, sessionId, roleId)
}

// (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
// or not to perform a given operation on a given object
func (sessionObject *SessionObject) CheckAccess(sessionId, operationId, objectId int) (bool, error) {
    return sessionObject.DBService.CheckAccess(sessionId, operationId, objectId)
}
