package RBAC

import (
    "testing"
)

func TestAddUser(t *testing.T) {
    // Add user - what we are actually testing
    setupUserTest()
    user, err := userObject.AddUser("test-user")
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = userObject.DeleteUser(user.Id)
    tearDownUserTest()
}

func TestDeleteUser(t *testing.T) {
    // Add user - what we are actually testing
    setupUserTest()
    user, err := userObject.AddUser("test-user")
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = userObject.DeleteUser(user.Id)
    tearDownUserTest()
}

func TestAssignedRoles(t *testing.T) {
    user := User{Id: 1}

    _, err := AssignedRoles(user.Id)

    if err != nil {
        t.Errorf("%v", err)
    }
}
