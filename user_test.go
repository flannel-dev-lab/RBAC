package RBAC

import (
    "testing"
)


func TestAddUser(t *testing.T) {
    // Add user - what we are actually testing
    user, err := AddUser("test-user")
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = DeleteUser(user.Id)
}

func TestDeleteUser(t *testing.T) {
    // Add user - what we are actually testing
    user, err := AddUser("test-user")
    if err != nil {
        t.Errorf("%v", err)
    }
    
    // Delete user - what we are actually testing
    _, err = DeleteUser(user.Id)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestAssignedRoles(t *testing.T) {
    user := User{Id: 1}

    _, err := AssignedRoles(user.Id)

    if err != nil {
        t.Errorf("%v", err)
    }
}
