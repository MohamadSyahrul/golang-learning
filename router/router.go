package router

import (
	controller "hacktiv8-go/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	bookRouter := router.Group("api/book")
	{
		bookRouter.GET("/get-all", controller.GetBookAll)
		bookRouter.GET("/by-id/:id", controller.GetBookById)
		bookRouter.POST("/create", controller.CreateBook)
		bookRouter.PUT("/update/:id", controller.UpdateBook)
		bookRouter.DELETE("/delete/:id", controller.DeleteBook)
	}

	return router

}
