package migration

import (
	"gorm.io/gorm"

	"github.com/tulioguaraldo/online-bookstore/model"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&model.Book{},
	)
}
