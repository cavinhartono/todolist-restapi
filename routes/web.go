package routes

import (
	"github.com/cavinhartono/todolist-restapi/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.default()

	bookgr := router.Group("/book")
	{
		router.GET("/products", controllers.GetAllBooks())
		router.GET("/products/:id", controllers.GetBookByID())
		router.POST("/products", controllers.CreateBook())
		router.PUT("/products/:id", controllers.UpdateBook())
		router.DELETE("/products/:id", controllers.DeleteBook())
	}
	return router
}
