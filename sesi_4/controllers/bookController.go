package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sesi_4_project/models"
)

type CreateOrUpdateBookInput struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}

func AllBooks(ctx *gin.Context) {
	var BookDatas []models.Book
	db := ctx.MustGet("db").(*gorm.DB)
	db.Find(&BookDatas)
	ctx.JSON(http.StatusOK, BookDatas)
}

func GetByBookId(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	bookId := ctx.Param("bookId")
	var bookData models.Book

	err := db.Where("id = ?", bookId).First(&bookData).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookId),
		})
		return
	}
	ctx.JSON(http.StatusOK, bookData)
}

func CreateBook(ctx *gin.Context) {
	var newBook CreateOrUpdateBookInput
	db := ctx.MustGet("db").(*gorm.DB)
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	book := models.Book{
		NameBook: newBook.NameBook,
		Author:   newBook.Author,
	}

	err := db.Create(&book).Error

	if err != nil {
		fmt.Println("Error creating book data:", err)
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

func UpdateBookById(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	db := ctx.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", bookId).First(&book).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	var input CreateOrUpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInputBook models.Book
	updatedInputBook.NameBook = input.NameBook
	updatedInputBook.Author = input.Author

	err := db.Model(&book).Where("id = ?", bookId).Updates(&updatedInputBook).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"errro_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func DeleteBookById(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	bookId := ctx.Param("bookId")
	book := models.Book{}

	var bookWhereId models.Book
	if err := db.Where("id = ?", bookId).First(&bookWhereId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	err := db.Where("id = ?", bookId).Delete(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"errro_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
