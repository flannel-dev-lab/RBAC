package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/vars"
    "testing"
)

var TestObject = vars.Object{Id: 2, Name: "test-fluid-object-name", Description: "Reserved object for fluid testing"}


func TestRoleOperationsOnObject(t *testing.T) {
    DbConnect("mysql", "asdf", "asdfaasdf", "awefawef", "rerg", 3306)


    role := vars.Role{Id: 1, Name: "test-name", Description: "Test Description"}
    object := vars.Object{Id: 1, Name: "test-object", Description: "Test Description"}
    _, err := RoleOperationsOnObject(role, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestUserOperationsOnObject(t *testing.T) {
    user := vars.User{Id: 1}
    object := vars.Object{Id: 1, Name: "objectName", Description: "Reserved object for testing"}
    _, err := UserOperationsOnObject(user, object)

    if err != nil {
        t.Errorf("%v", err)
    }
}

func TestCreateObject(t *testing.T) {
    setupRBACObjectTest()
    object, err := rbacObject.CreateObject("test-object", "test-object-description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = rbacObject.RemoveObject(object.Id)
    tearDownRBACObjectTest()
}

func TestRemoveObject(t *testing.T) {
    setupRBACObjectTest()
    // Create an object to remove
    object, err := rbacObject.CreateObject("test-object", "test-object-description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Remove the object
    _, err = rbacObject.RemoveObject(object.Id)

    if err != nil {
        t.Errorf("%v", err)
    }
    tearDownRBACObjectTest()
}

