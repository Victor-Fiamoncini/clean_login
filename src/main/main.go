package main

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/main/config"

func main() {
	app := config.NewApp()

	app.Listen(":3333")
}
