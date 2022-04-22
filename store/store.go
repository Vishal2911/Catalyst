package store

import (
	"magantas/catalyst/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db    *gorm.DB
	Token map[string]model.TokenSchema
}



func NewDB() *Store {
	// Openning file
	db, err := gorm.Open("sqlite3", "./catalyst.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	// Creating the table
	if !db.HasTable(&model.Users{}) {
		db.CreateTable(&model.Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.Users{})
	}
	tokens := make(map[string]model.TokenSchema)
	return &Store{
		db: db,
		Token: tokens,
	}
}
