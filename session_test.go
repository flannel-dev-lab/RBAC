package RBAC

import (
    "testing"
)


func TestCreateSession(t *testing.T) {
    user := User{UserId: 1}
    session := Session{SessionId: 1, UserId: 1, Token: "123-123-123"}

    _, err := CreateSession(user, session)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDeleteSession(t *testing.T) {
    user := User{UserId: 1}

    _, err := DeleteSession(user, "testSession")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestAddActiveRole(t *testing.T) {
    user := User{UserId: 1}
    session := Session{SessionId: 1, UserId: 1, Token: "123-123-123"}

    _, err := AddActiveRole(user, session, "testRole")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDropActiveRole(t *testing.T) {
    user := User{UserId: 1}
    session := Session{SessionId: 1, UserId: 1, Token: "123-123-123"}

    _, err := DropActiveRole(user, session, "testRole")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestCheckAccess(t *testing.T) {
    session := Session{SessionId: 1, UserId: 1, Token: "123-123-123"}
    operation := Operation{OperationId: 1, Name: "testOperation", Description: "Reserved permission for test"}
    object := Object{ObjectId: 1, Name: "testObject", Description: "Reserved object for testing"}
    
    _, err := CheckAccess(session, operation, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}
