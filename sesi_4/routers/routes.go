package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sesi_4_project/controllers"
)

func LoadRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/books", controllers.AllBooks)
	r.GET("/books/:bookId", controllers.GetByBookId)
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:bookId", controllers.UpdateBookById)
	r.DELETE("/books/:bookId", controllers.DeleteBookById)

	return r
}
