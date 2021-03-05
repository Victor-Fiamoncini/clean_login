package main

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/bootstrap/config"

func main() {
	app := config.NewApp()

	app.Listen(":3001")
}
