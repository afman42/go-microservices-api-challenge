package routers

import (
	"database/sql"
	"sesi_3_challenge/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.GET("/book", controllers.AllBooks)
	router.GET("/book/:bookID", controllers.GetBook)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
