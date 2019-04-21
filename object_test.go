package RBAC

import ( 
    "testing"
)

var TestObject = Object{Id: 2, Name: "test-fluid-object-name", Description: "Reserved object for fluid testing"}


func TestRoleOperationsOnObject(t *testing.T) {
    _, err := RoleOperationsOnObject(TestRoleStatic, TestObject)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserOperationsOnObject(t *testing.T) {
    user := User{UserId: 1}
    object := Object{Id: 1, Name: "objectName", Description: "Reserved object for testing"}
    _, err := UserOperationsOnObject(user, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestCreateObject(t *testing.T) {
    object, err := CreateObject("test-object", "test-object-description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = RemoveObject(object)
}

func TestRemoveObject(t *testing.T) {
    // Create an object to remove
    object, err := CreateObject("test-object", "test-object-description")

    // Remove the object
    _, err = RemoveObject(object)

    if err != nil {
        t.Errorf("%v", err)
    }
}
