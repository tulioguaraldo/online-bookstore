package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tulioguaraldo/online-bookstore/db/migration"
)

var db *gorm.DB

func StartDb() {
	database, err := gorm.Open(sqlite.Open("../sqlite/store.db"), &gorm.Config{})
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect to Sqlite %s", err.Error())
		log.Fatal(errMessage)
	}

	fmt.Println("Connected to Sqlite")

	db = database

	migration.Migration(db)
}

func ConnectToDb() *gorm.DB {
	StartDb()

	return db
}
