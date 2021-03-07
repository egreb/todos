package database

import "github.com/go-pg/pg/v10"

// Connect to the postgres database
func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "sib",
		Database: "todos",
		Password: "",
	})

	return db
}
