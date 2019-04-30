package rbac

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
	"log"
	"os"
	"testing"
)

func setupUserTest(userObject *UserObject) {
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

func tearDownUserTest(userObject *UserObject) {
	err := userObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestAddUser(t *testing.T) {
	// Add user - what we are actually testing
	var userObject UserObject
	setupUserTest(&userObject)
	user, err := userObject.AddUser("test-user")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Cleanup
	_, err = userObject.DeleteUser(user.Id)
	tearDownUserTest(&userObject)
}

func TestDeleteUser(t *testing.T) {
	// Add user - what we are actually testing
	var userObject UserObject
	setupUserTest(&userObject)
	user, err := userObject.AddUser("test-user")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Cleanup
	_, err = userObject.DeleteUser(user.Id)
	tearDownUserTest(&userObject)
}

func TestAssignedRoles(t *testing.T) {
	var userObject UserObject
	setupUserTest(&userObject)
	user := vars.User{Id: 1}

	_, err := userObject.AssignedRoles(user.Id)

	if err != nil {
		t.Errorf("%v", err)
	}
	tearDownUserTest(&userObject)
}
