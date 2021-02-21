package config

import (
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv func
func LoadEnv() {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	rootDir := filepath.Dir(wd)

	err = godotenv.Load(path.Join(rootDir, "../.env"))

	if err != nil {
		panic(err.Error())
	}
}
