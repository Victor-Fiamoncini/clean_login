package config

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"

// CreateSchemas func
func CreateSchemas() {
	db := database.OpenConnection()

	err := database.CreateSchemas(db)

	if err != nil {
		panic(err.Error())
	}
}
