package RBAC

import (
    "testing"
)


func TestAddUser(t *testing.T) {
    _, err := AddUser("testUser")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDeleteUser(t *testing.T) {
    _, err := DeleteUser("testUser")

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestAssignedRoles(t *testing.T) {
    user := User{UserId: 1}

    _, err := AssignedRoles(user)

    if err != nil {
        t.Errorf("%v", err)
    }
}
