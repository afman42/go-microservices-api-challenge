package routers

import (
	"sesi_4_project/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/book", controllers.AllBooks)
	r.GET("/book/:bookId", controllers.GetByBookId)
	r.POST("/book", controllers.CreateBook)
	r.PUT("/book/:bookId", controllers.UpdateBookById)
	r.DELETE("/book/:bookId", controllers.DeleteBookById)

	return r
}
