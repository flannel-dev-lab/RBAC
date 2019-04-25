package RBAC

import (
	"github.com/flannel-dev-lab/RBAC/database"
	"os"
)

var (
	DbReady bool
	DB      database.DatabaseService
)

// Used for testing to load DB from environment variables
func DbInit() {
	if !DbReady {
		driver := os.Getenv("RBAC_DB_DRIVER")
		username := os.Getenv("RBAC_DB_USERNAME")
		password := os.Getenv("RBAC_DB_PASSWORD")
		hostname := os.Getenv("RBAC_DB_HOSTNAME")
		dbname := os.Getenv("RBAC_DB_NAME")
		port := os.Getenv("RBAC_DB_PORT")

		DbConnect(driver, username, password, hostname, dbname, port)

		DbReady = true
	}
}

func DbConnect(driver string, username string, password string, hostname string, databaseName string, port string) {
	DB, _ := database.CreateDatabaseObject(driver)
	DB.CreateDBConnection(driver, username, password, hostname, databaseName, port)
}
