package database

import (
	"egreb.net/todos/todo"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// CreateSchema run migrations
func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*todo.Entity)(nil),
	}

	for _, m := range models {
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
