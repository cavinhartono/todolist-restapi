package controllers

import (
	"fmt"
	"net/http"

	"todolist-restapi/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.CreateBook(&book)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.GetBookByID(&book, id)
	if err != nil {
		c.JSON(http.StatusNotFound, book)
	}
	c.BindJSON(&book)
	err = models.UpdateBook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Params.ByName("id")
	err := models.DeleteBook(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func GetAllBooks(c *gin.Context) {
	var book []models.Book
	err := models.GetAllBooks(&book)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}

func GetBookByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var book models.Book
	err := models.GetBookByID(&book, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, book)
	}
}
