package routers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"sesi_3_challenge/controllers"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.GET("/books", controllers.AllBooks)
	router.GET("/book/:bookID", controllers.GetBook)
	router.POST("/book/add", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
