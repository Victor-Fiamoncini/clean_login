package config

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"

// CreateSchemas func
func CreateSchemas() {
	connection := database.OpenConnection("localhost", "postgres", "postgres", "5432", "auth")

	err := database.CreateSchemas(connection)

	if err != nil {
		panic(err.Error())
	}
}
