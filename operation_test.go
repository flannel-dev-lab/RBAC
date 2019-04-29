package RBAC

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"log"
	"os"
	"testing"
)

// RBAC Object test parameters
var operationObject OperationObject

func setupOperationObjectTest() {
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

	operationObject.DBService = dbService
}

func tearDownOperationObjectTest() {
	err := operationObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestAddOperation(t *testing.T) {
	setupOperationObjectTest()
	object, err := operationObject.AddOperation("test-operation", "test-operation-description")

	if err != nil {
		t.Errorf("%v", err)
	}

	// Cleanup
	_, err = operationObject.DeleteOperation(object.Id)
	tearDownOperationObjectTest()
}

func TestDeleteOperation(t *testing.T) {
	setupOperationObjectTest()
	object, err := operationObject.AddOperation("test-operation", "test-operation-description")

	if err != nil {
		t.Errorf("%v", err)
	}

	// Cleanup
	_, err = operationObject.DeleteOperation(object.Id)
	tearDownOperationObjectTest()
}

