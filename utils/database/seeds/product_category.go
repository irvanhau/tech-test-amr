package seeds

import (
	"TechnicalTest/features/products"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateProductCategory(db *gorm.DB, name string) error {

	var countData int64
	db.Table("product_categories").Count(&countData)

	if countData < 5 {
		return db.Create(products.ProductCategory{
			ID:   uuid.New(),
			Name: name,
		}).Error
	}

	return nil
}
