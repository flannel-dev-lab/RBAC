package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
    "log"
    "os"
    "testing"
)


func setupRBACObjectTest(rbacObject *RBACObject) {
    dbService, err := database.CreateDatabaseObject("mysql")
    if err != nil {
        log.Fatalf(err.Error())
    }

    err = dbService.CreateDBConnection(
        os.Getenv("RBAC_DB_DRIVER"),
        os.Getenv("RBAC_DB_USERNAME"),
        os.Getenv("RBAC_DB_PASSWORD"),
        os.Getenv("RBAC_DB_HOSTNAME"),
        os.Getenv("RBAC_DB_NAME"),
        os.Getenv("RBAC_DB_PORT"))

    if err != nil {
        log.Fatalf(err.Error())
    }

    rbacObject.DBService = dbService
}

func tearDownRBACObjectTest(rbacObject *RBACObject) {
    err := rbacObject.DBService.CloseConnection()
    if err != nil {
        log.Fatalf(err.Error())
    }
}

func TestCreateObject(t *testing.T) {
    var rbacObject RBACObject

    setupRBACObjectTest(&rbacObject)
    object, err := rbacObject.CreateObject("test-object", "test-object-description")

    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = rbacObject.RemoveObject(object.Id)
    tearDownRBACObjectTest(&rbacObject)
}

func TestRemoveObject(t *testing.T) {
    var rbacObject RBACObject
    setupRBACObjectTest(&rbacObject)
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
    tearDownRBACObjectTest(&rbacObject)
}

