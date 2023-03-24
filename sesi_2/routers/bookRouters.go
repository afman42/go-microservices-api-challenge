package routers

import (
	"challenge_sesi_2_api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/book", controllers.AllBooks)
	router.GET("/book/:bookID", controllers.GetBook)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
