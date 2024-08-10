package products

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  string    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductCategory struct {
	ID   uuid.UUID `json:"category_id"`
	Name string    `json:"category_name"`
}

type ProductInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type ProductHandlerInterface interface {
	GetProducts() echo.HandlerFunc
	CreateProduct() echo.HandlerFunc
}

type ProductServiceInterface interface {
	GetProducts(sort, keyword string, filter string, offset int) ([]ProductInfo, map[string]any, error)
	CreateProduct(NewData Product) (*Product, error)
}

type ProductDataInterface interface {
	GetAll(sort, keyword string, filter string, offset int) ([]ProductInfo, map[string]any, error)
	Insert(newData Product) (*Product, error)
}
