package RBAC

import (
    "errors"
)

// A Session represents a user as owner and an active role set
type Session struct {
    SessionId       int
    UserId          int
    Token           string
}

// (RC-16) Core RBAC: Create a new session with a user as owner and an active role set
func CreateSession(user User, session Session) (bool, error) { 
    return false, errors.New("Not yet implemented")
}

// (RC-23) Core RBAC: Delete a given session with a given owner user
func DeleteSession(user User, sessionName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-01) Core RBAC: Add a role as an active role of a session whose owner is a given user
func AddActiveRole(user User, session Session, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-27) Core RBAC: Delete a role from the active role set of a session owned by a given user
func DropActiveRole(user User, session Session, roleName string) (bool, error) {
    return false, errors.New("Not yet implemented")
}

// (RC-14) Core RBAC: Returns a boolean of whether the subject of a given session is allowed
// or not to perform a given operation on a given object
func CheckAccess(session Session, operation Operation, object Object) (bool, error) {
    return false, errors.New("Not yet implemented")
}
