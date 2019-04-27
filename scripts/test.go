package main

import (
	"fmt"
	"github.com/flannel-dev-lab/RBAC"
	"github.com/flannel-dev-lab/RBAC/database"
)

func main()  {
	dbService, err  := database.CreateDatabaseObject("mysql")

	if err != nil {
		fmt.Println(err)
	}

	err = dbService.CreateDBConnection("mysql", "root", "changemeyombu", "dev-applayer-mariadb.c9iyojqezmkr.us-east-1.rds.amazonaws.com", "dev", "3306")

	if err != nil {
		fmt.Println(err)
	}

	userObject :=  RBAC.UserObject{}
	roleObject := RBAC.RoleObject{}
	sessionObject := RBAC.SessionObject{}

	userObject.DBService = dbService
	roleObject.DBService = dbService
	sessionObject.DBService = dbService

	fmt.Println(roleObject.DeleteRole(111))
	fmt.Println(userObject.DeleteUser(110))



}
