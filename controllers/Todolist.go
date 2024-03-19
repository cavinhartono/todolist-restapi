package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cavinhartono/models"
)

var products = []Product {
	{ ID: 1, Name: "Vans Oldschool", Price: 600 },
	{ ID: 2, Name: "Adidas Gazelle", Price: 700 },
	{ ID: 3, Name: "Cosmic Green Shirt", Price: 120 },
}

func createProducts(c *gin.Context) {
	var newProduct Product
	
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}
	
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func updateProducts(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product
		
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{ "error": "Product not found" })
}

func deleteProducts(c *gin.Context) {
	id := c.Param("id")
	
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{ "message": "Product is deleted" })
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{ "error": "Product not found" })
}

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func getProducts(c *gin.Context) {
	id := c.Param("id")
	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{ "error": "Product not found" })
}