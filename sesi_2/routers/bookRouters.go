package routers

import (
	"challenge_sesi_2_api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.AllBooks)
	router.GET("/book/:bookID", controllers.GetBook)
	router.POST("/book/add", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
