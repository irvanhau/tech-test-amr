package seeds

import (
	"TechnicalTest/utils/database/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "Create Product Category 1",
			Run:  func(db *gorm.DB) error { return CreateProductCategory(db, "Minuman") },
		},
		{
			Name: "Create Product Category 2",
			Run:  func(db *gorm.DB) error { return CreateProductCategory(db, "Makanan") },
		},
		{
			Name: "Create Product Category 3",
			Run:  func(db *gorm.DB) error { return CreateProductCategory(db, "Elektronik") },
		},
		{
			Name: "Create Product Category 4",
			Run:  func(db *gorm.DB) error { return CreateProductCategory(db, "Handphone") },
		},
		{
			Name: "Create Product Category 5",
			Run:  func(db *gorm.DB) error { return CreateProductCategory(db, "Kendaraan") },
		},
	}
	return seeds
}
