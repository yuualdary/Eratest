package main

import (
	"eratest/Product"
	"eratest/config"
	"eratest/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.ConnectDatabase()

	ProductRepository := Product.NewRepository(config.DB)
	ProductService := Product.NewService(ProductRepository)
	ProductHandler := handler.NewProductHandler(ProductService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/product", ProductHandler.CreateProduct)
		v1.GET("/product", ProductHandler.NewProduct)

	}

	router.Run(":8000")

}
