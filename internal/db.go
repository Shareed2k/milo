package internal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(settings Settings) *Database {
	db, err := gorm.Open("sqlite3", "test.db?cache=shared&mode=rwc")

	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}

func (d *Database) Close() {
	d.DB.Close()
}
