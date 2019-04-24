package RBAC

import (
    "testing"
)


var (
    TestRoleStatic = Role{Id:1, Name: "test-fluid-role-name", Description: "Reserved role for fluid testing"}
)


func TestAddRole(t *testing.T) {
    role, err := AddRole("test-role", "Test role description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = DeleteRole(role.Id)
}

func TestDeleteRole(t *testing.T) {
    // Create a role to delete
    role, err := AddRole("test-role", "Test role description")
    if err != nil {
        t.Errorf("%v", err)
    }

    // Delete the role - this is what we are really testing
    _, err = DeleteRole(role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestAssignUser(t *testing.T) {
    // Create a new user
    user, userErr := AddUser("test-user")
    if userErr != nil {
        t.Errorf("%v", userErr)
    }

    // Create a new role
    role, roleErr := AddRole("test-role", "Test role description")
    if roleErr != nil {
        t.Errorf("%v", roleErr)
    }

    // Assign the user - what we are actually testing
    _, err := AssignUser(user, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = DeassignUser(user, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = DeleteUser(user.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = DeleteRole(role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDeassignUser(t *testing.T) {
    // Create a new user
    user, userErr := AddUser("test-user")
    if userErr != nil {
        t.Errorf("%v", userErr)
    }

    // Create a new role
    role, roleErr := AddRole("test-role", "Test role description")
    if roleErr != nil {
        t.Errorf("%v", roleErr)
    }

    // Assign the user
    _, err := AssignUser(user, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Deassign the user - what we are actually testing
    _, err = DeassignUser(user, role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = DeleteUser(user.Id)
    if err != nil {
        t.Errorf("%v", err)
    }

    _, err = DeleteRole(role.Id)
    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestAssignedUsers(t *testing.T) {
    role := Role{Id: 1, Name: "roleName", Description:"Reserved role for testing"}
    _, err := AssignedUsers(role.Id);

    if err != nil {
        t.Errorf("%v", err);
    }
}

func TestSessionRoles(t *testing.T) {
    session := Session{Id: 1, UserId: 1, Name: "123-123-123"}
    _, err := SessionRoles(session);

    if err != nil {
        t.Errorf("%v", err);
    }
}
