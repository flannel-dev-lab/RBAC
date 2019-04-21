package RBAC

import (
    "testing"
)

const TestUserNameFluid = "test-user-name-fluid"

func TestAddUser(t *testing.T) {
    _, err := AddUser(TestUserNameFluid)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestDeleteUser(t *testing.T) {
    _, err := DeleteUser(TestUserNameFluid)

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
