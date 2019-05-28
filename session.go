package rbac

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
)

// SessionObject Manages the sessions
type SessionObject struct {
	DBService database.DatabaseService
}

// CreateSession (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
func (sessionObject *SessionObject) CreateSession(userId int, sessionName string) (vars.Session, error) {
	return sessionObject.DBService.CreateSession(userId, sessionName)
}

// DeleteSession (RC-23) Core RBAC: Delete a given session with a given owner user
func (sessionObject *SessionObject) DeleteSession(userId int, sessionName string) (bool, error) {
	return sessionObject.DBService.DeleteSession(userId, sessionName)
}

// AddActiveRole (RC-01) Core RBAC: Add a role as an active role of a session whose owner is a given user
func (sessionObject *SessionObject) AddActiveRole(userId int, sessionName, roleName string) (bool, error) {
	return sessionObject.DBService.AddActiveRole(userId, sessionName, roleName)
}

// DropActiveRole (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
func (sessionObject *SessionObject) DropActiveRole(userId int, sessionName, roleName string) (bool, error) {
	return sessionObject.DBService.DropActiveRole(userId, sessionName, roleName)
}

// CheckAccess (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
// or not to perform a given operation on a given object
func (sessionObject *SessionObject) CheckAccess(sessionName, operationName, objectName string) (bool, error) {
	return sessionObject.DBService.CheckAccess(sessionName, operationName, objectName)
}
