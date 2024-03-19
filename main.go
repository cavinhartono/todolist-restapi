package main

import (
	"github.com/cavinhartono/controllers"
	"github.com/cavinhartono/models"
)

func main()  {
	router := gin.default()

	router.GET("/products", getAllProducts)
	router.GET("/products/:id", getProducts)
	router.POST("/products", createProducts)
	router.PUT("/products/:id", updateProducts)
	router.DELETE("/products/:id", deleteProducts)

	router.Run(":8080")
}