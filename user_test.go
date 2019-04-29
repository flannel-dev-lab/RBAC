package RBAC

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"github.com/flannel-dev-lab/RBAC/vars"
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

func tearDownUserTest() {
	err := userObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestAddUser(t *testing.T) {
	// Add user - what we are actually testing
	setupUserTest()
	_, err := userObject.AddUser("test-user")
	if err != nil {
		t.Errorf("%v", err)
	}

	// Cleanup
	//_, err = userObject.DeleteUser(user.Id)
	//tearDownUserTest()
}

func TestDeleteUser(t *testing.T) {
	// Add user - what we are actually testing
	setupUserTest()
	user, err := userObject.AddUser("test-user")
	if err != nil {
		t.Errorf("%v", err)
	}

	// Cleanup
	_, err = userObject.DeleteUser(user.Id)
	tearDownUserTest()
}

func TestAssignedRoles(t *testing.T) {
	setupUserTest()
	user := vars.User{Id: 1}

	_, err := userObject.AssignedRoles(user.Id)

	if err != nil {
		t.Errorf("%v", err)
	}
	tearDownUserTest()
}
