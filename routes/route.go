package routes

import (
	"TechnicalTest/features/products"
	"github.com/labstack/echo/v4"
)

func RouteProduct(e *echo.Echo, ph products.ProductHandlerInterface) {
	e.GET("/product", ph.GetProducts())
	e.POST("/product", ph.CreateProduct())
}
