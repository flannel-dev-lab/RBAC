package RBAC

import (
    "github.com/flannel-dev-lab/RBAC/database"
    "log"
    "os"
    "testing"
)

var userObject UserObject

func setupUserTest() {
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

    userObject.DBService = dbService
}

func TestAddUser(t *testing.T) {
    // Add user - what we are actually testing
    err := userObject.AddUser("test-user")
    if err != nil {
        t.Errorf("%v", err)
    }

    // Cleanup
    _, err = userObject.DeleteUser(userObject.User.Id)
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
