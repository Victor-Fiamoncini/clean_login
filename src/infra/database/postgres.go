package database

import (
	"fmt"
	"sync"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var connectionInstance *pg.DB
var once sync.Once

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	port     = "5432"
	dbName   = "auth"
)

// OpenConnection func
func OpenConnection() *pg.DB {
	if connectionInstance == nil {
		once.Do(func() {
			connectionInstance = pg.Connect(&pg.Options{
				User:     user,
				Password: password,
				Addr:     fmt.Sprintf("%s:%s", host, port),
				Database: dbName,
			})
		})
	}

	return connectionInstance
}

// CloseConnection func
func CloseConnection(db *pg.DB) error {
	err := db.Close()

	if err != nil {
		return err
	}

	return nil
}

// CreateSchemas func
func CreateSchemas(db *pg.DB) error {
	models := []interface{}{
		(*entities.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
