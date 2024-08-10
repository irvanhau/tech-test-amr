package database

import (
	"TechnicalTest/features/products/data"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(data.ProductCategory{})
	db.AutoMigrate(data.Product{})
}
