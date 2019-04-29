package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
    "log"
    "os"
    "testing"
)

// RBAC Object test parameters
var rbacObject RBACObject

func setupRBACObjectTest() {
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

func tearDownRBACObjectTest() {
    err := rbacObject.DBService.CloseConnection()
    if err != nil {
        log.Fatalf(err.Error())
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

