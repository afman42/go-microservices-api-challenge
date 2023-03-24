package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sesi_3_challenge/models"

	"github.com/gin-gonic/gin"
)

type CreateOrUpdateBookInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func AllBooks(ctx *gin.Context) {
	var results = []models.Book{}

	sqlStatement := "SELECT id, title, author, description FROM books"

	rows, err := ctx.MustGet("db").(*sql.DB).Query(sqlStatement)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		results = append(results, book)
	}

	if err = rows.Err(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func GetBook(ctx *gin.Context) {
	var bookID = ctx.Param("bookID")
	var db = ctx.MustGet("db").(*sql.DB)
	var bookData = models.Book{}
	var sqlStatement = "SELECT * FROM books WHERE id = $1"

	err := db.QueryRow(sqlStatement, bookID).
		Scan(&bookData.ID, &bookData.Title, &bookData.Author, &bookData.Description)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, bookData)
}

func CreateBook(ctx *gin.Context) {
	var newBook CreateOrUpdateBookInput
	var book = models.Book{}
	var db = ctx.MustGet("db").(*sql.DB)

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	sqlStatement := `
		INSERT INTO books (title,author,description)
		VALUES ($1,$2,$3)
		RETURNING *
	`
	err := db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Description).
		Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

func UpdateBook(ctx *gin.Context) {
	var bookID = ctx.Param("bookID")
	var db = ctx.MustGet("db").(*sql.DB)
	var book models.Book
	var sqlStatement = `
		UPDATE books
		SET title = $2, author = $3, description = $4
		WHERE id = $1;
	`
	var idWhere int
	var sqlStatementSelectId = `
		SELECT id FROM books WHERE id = $1
	`
	err := db.QueryRow(sqlStatementSelectId, bookID).
		Scan(&idWhere)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	var input CreateOrUpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	book.ID = bookID
	book.Title = input.Title
	book.Author = input.Author
	book.Description = input.Description

	res, err := db.Exec(sqlStatement, book.ID, book.Title, book.Author, book.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, err := res.RowsAffected()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	if count > 0 {
		ctx.JSON(http.StatusOK, book)
	}
}

func DeleteBook(ctx *gin.Context) {
	var bookID = ctx.Param("bookID")
	var db = ctx.MustGet("db").(*sql.DB)
	var idWhere int
	var sqlStatementSelectId = `
		SELECT id FROM books WHERE id = $1
	`
	err := db.QueryRow(sqlStatementSelectId, bookID).
		Scan(&idWhere)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})
		return
	}

	var sqlStatement = `
		DELETE FROM books
		WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfull deleted", bookID),
	})
}
