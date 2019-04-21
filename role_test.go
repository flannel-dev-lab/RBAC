package RBAC

import (
    "testing"
)

const TestRoleNameFluid = "test-role-name-fluid"
const TestRoleDescriptionFluid = "Reserved for fluid role description testing"

var (
    TestRoleStatic = Role{Id:1, Name: "test-fluid-role-name", Description: "Reserved role for fluid testing"}
)


func TestAddRole(t *testing.T) {
    _, err := AddRole(TestRoleNameFluid, TestRoleDescriptionFluid)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDeleteRole(t *testing.T) {
    _, err := DeleteRole(TestRoleNameFluid)
 
    if err != nil {
        t.Errorf("%v", err);
    }
}

func TestAssignUser(t *testing.T) {
    user := User{UserId: 1}
    _, err := AssignUser(user, "roleName");

    if err != nil {
        t.Errorf("%v", err);
    }
}

func TestDeassignUser(t *testing.T) {
    user := User{UserId: 1}
    _, err := DeassignUser(user, "roleName");

    if err != nil {
        t.Errorf("%v", err);
    }
}

func TestAssignedUsers(t *testing.T) {
    role := Role{Id: 1, Name: "roleName", Description:"Reserved role for testing"}
    _, err := AssignedUsers(role);

    if err != nil {
        t.Errorf("%v", err);
    }
}

func TestSessionRoles(t *testing.T) {
    session := Session{SessionId: 1, UserId: 1, Token: "123-123-123"}
    _, err := SessionRoles(session);

    if err != nil {
        t.Errorf("%v", err);
    }
}
