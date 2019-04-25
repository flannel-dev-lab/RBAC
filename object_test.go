package RBAC

import ( 
    "testing"
)

var TestObject = Object{Id: 2, Name: "test-fluid-object-name", Description: "Reserved object for fluid testing"}


func TestRoleOperationsOnObject(t *testing.T) {
    DbConnect("mysql", "asdf", "asdfaasdf", "awefawef", "rerg", 3306)


    role := Role{Id: 1, Name: "test-name", Description: "Test Description"}
    object := Object{Id: 1, Name: "test-object", Description: "Test Description"}
    _, err := RoleOperationsOnObject(role, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserOperationsOnObject(t *testing.T) {
    user := User{Id: 1}
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
