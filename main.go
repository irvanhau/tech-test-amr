package main

import (
	"TechnicalTest/configs"
	"TechnicalTest/features/products/data"
	"TechnicalTest/features/products/handler"
	"TechnicalTest/features/products/service"
	"TechnicalTest/helpers"
	"TechnicalTest/helpers/generate_uuid"
	"TechnicalTest/routes"
	"TechnicalTest/utils/cache"
	"TechnicalTest/utils/database"
	"TechnicalTest/utils/database/seeds"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType},
	}))

	config := configs.InitConfig()
	db := database.InitDB(config)
	database.Migrate(db)

	redis := cache.InitRedis()

	genUUID := generate_uuid.InitUUID()
	time := helpers.InitTime()

	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			fmt.Printf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	productData := data.New(db, redis)
	productService := service.New(productData, genUUID, time)
	productHandler := handler.NewHandler(productService)

	routes.RouteProduct(e, productHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Server)).Error())
}
