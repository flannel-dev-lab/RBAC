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
    _, err = RemoveObject(object)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestRevokePermission(t *testing.T) {
    // Create an object
    object, err := CreateObject("test-object", "test-object-description")
    if err != nil {
        t.Errorf("%v", err)
    }
    
    // Grant a permission
    _, err = GrantPermission(object, TestOperation, 1)
    if err != nil {
        t.Errorf("%v", err)
    }

    // This is what we are actually testing for
    _, err = RevokePermission(object, TestOperation, 1)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = RemoveObject(object)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestRolePermissions(t *testing.T) {
    role := Role{Id: 1, Name: "test-role", Description: "Reserved role for testing"}
    _, err := RolePermissions(role)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserPermissions(t *testing.T) {
    user := User{Id: 1}

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
