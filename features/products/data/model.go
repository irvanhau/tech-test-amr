package data

import (
	"github.com/google/uuid"
	"time"
)

type ProductCategory struct {
	ID      uuid.UUID `gorm:"column:id;unique;type:uuid;primary_key"`
	Name    string    `gorm:"column:name;type:varchar(255);not null"`
	Product []Product `gorm:"foreignkey:CategoryID"`
}

type Product struct {
	ID          uuid.UUID `gorm:"column:id;unique;type:uuid;primary_key"`
	CategoryID  string    `gorm:"column:category_id;type:uuid;not null"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:text;not null"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
