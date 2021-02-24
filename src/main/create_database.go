package main

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"

func init() {
	connection := database.OpenConnection(
		"localhost",
		"postgres",
		"victor",
		5432,
		"auth_clean_architecture",
	)

	database.CreateSchemas(connection)
}
