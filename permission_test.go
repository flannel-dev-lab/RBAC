package RBAC

import (
    "testing"
)


func TestGrantPermission(t *testing.T) {
    // Create an object
    object, err := CreateObject("test-object", "test-object-description")

    _, err = GrantPermission(object, TestOperation, 1)

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    //_, err = RemoveObject(object)
}

func TestRevokePermission(t *testing.T) {
    object := Object{Id: 1, Name: "testObject", Description: "Reserved object for testing"}
    operation := Operation{Id: 1, Name: "testOperation", Description: "Reserved permission for test"}
    _, err := RevokePermission(object, operation, "testRole")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestRolePermissions(t *testing.T) {
    role := Role{Id: 1, Name: "roleName", Description:" Reserved role for testing"}
    _, err := RolePermissions(role)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserPermissions(t *testing.T) {
    user := User{UserId: 1}

    _, err := UserPermissions(user)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestSessionPermissions(t *testing.T) {
    session := Session{SessionId: 1, UserId: 1, Token: "123-123-123"}
    _, err := SessionPermissions(session)

    if err != nil {
        t.Errorf("%v", err)
    }
}
