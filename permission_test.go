package rbac

import (
	"fmt"
	"github.com/flannel-dev-lab/RBAC/database"
	"log"
	"os"
	"testing"
)

func setupPermissionObjectTest(permissionObject *PermissionObject, rbacObject *RBACObject, operationObject *OperationObject, roleObject *RoleObject, userObject *UserObject, sessionObject *SessionObject) {
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

	permissionObject.DBService = dbService
	rbacObject.DBService = dbService
	operationObject.DBService = dbService
	roleObject.DBService = dbService
	userObject.DBService = dbService
	sessionObject.DBService = dbService
}

func tearDownPermissionObjectTest(permissionObject *PermissionObject, rbacObject *RBACObject, operationObject *OperationObject, roleObject *RoleObject, userObject *UserObject, sessionObject *SessionObject) {
	err := permissionObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = rbacObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = operationObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = roleObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = userObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = sessionObject.DBService.CloseConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestGrantPermission(t *testing.T) {
	var permissionObject PermissionObject
	var rbacObject RBACObject
	var operationObject OperationObject
	var roleObject RoleObject
	var userObject UserObject
	var sessionObject SessionObject

	setupPermissionObjectTest(&permissionObject, &rbacObject, &operationObject, &roleObject, &userObject, &sessionObject)

	object, err := rbacObject.CreateObject("test-object-1", "test-object-description")
	if err != nil {
		t.Errorf(err.Error())
	}

	operation, err := operationObject.AddOperation("test-operation-1", "test-operation-description")
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = permissionObject.CreatePermission(object.Id, operation.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	findPermission, err := permissionObject.FindPermission(object.Id, operation.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("FindPermission: ", findPermission)

	role, err := roleObject.AddRole("test-role", "test-role-description")
	if err != nil {
		t.Errorf(err.Error())
	}

	rolePermission, err := permissionObject.RolePermissions(role.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("RolePermission: ", rolePermission)

	user, err := userObject.AddUser("test-user")
	if err != nil {
		t.Errorf(err.Error())
	}

	userPermission, err := permissionObject.UserPermissions(user.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("UserPermission: ", userPermission)

	session, err := sessionObject.CreateSession(user.Id, "test-session")
	if err != nil {
		t.Errorf(err.Error())
	}

	sessionPermission, err := permissionObject.SessionPermissions(session.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("SessionPermission: ", sessionPermission)

	_, err = permissionObject.GrantPermission(object.Id, operation.Id, role.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = rbacObject.RemoveObject(object.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = operationObject.DeleteOperation(operation.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = roleObject.DeleteRole(role.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = userObject.DeleteUser(user.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = sessionObject.DeleteSession(user.Id, session.Id)
	if err != nil {
		t.Errorf(err.Error())
	}

	tearDownPermissionObjectTest(&permissionObject, &rbacObject, &operationObject, &roleObject, &userObject, &sessionObject)
}
