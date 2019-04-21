package RBAC

import ( 
    "testing"
)


func TestRoleOperationsOnObject(t *testing.T) {
    role := Role{RoleId: 1, Name: "roleName", Description:" Reserved role for testing"}
    object := Object{ObjectId: 1, Name: "objectName", Description: "Reserved object for testing"}

    _, err := RoleOperationsOnObject(role, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserOperationsOnObject(t *testing.T) {
    user := User{UserId: 1}
    object := Object{ObjectId: 1, Name: "objectName", Description: "Reserved object for testing"}
    _, err := UserOperationsOnObject(user, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}
