package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"mauthor"`
	Desc   string `json:"desc"`
}

var BookData = []Book{}

func GetBookAll(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"book": BookData,
	})

}

func GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for _, book := range BookData {
		if book.Id == id {
			ctx.JSON(http.StatusOK, book)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})

}

func CreateBook(ctx *gin.Context) {
	var reqBook Book

	if err := ctx.ShouldBindJSON(&reqBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqBook.Id = len(BookData) + 1
	BookData = append(BookData, reqBook)

	ctx.JSON(http.StatusCreated, reqBook)
}

func UpdateBook(ctx *gin.Context) {
	var request Book

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range BookData {
		if book.Id == id {
			request.Id = id
			BookData[i] = request
			ctx.JSON(http.StatusOK, request)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})

}

func DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	for i, book := range BookData {

		if book.Id == id {
			BookData = append(BookData[:i], BookData[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Book Deleted Success"})
			return
		}

	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
