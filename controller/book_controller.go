package controller

import (
	"awesomeProject/entity"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "select all books",
		"books":   service.GetAllBooks(),
	})
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}
	book, err := service.GetBookByID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "select book by id",
		"book":    book,
	})
}

func AddBook(c *gin.Context) {
	var book entity.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	book = service.InsertBook(book)
	c.JSON(http.StatusOK, gin.H{
		"message": "add book by id",
		"book":    book,
	})
}

func UpdateBookByID(c *gin.Context) {
	bookId, iderr := strconv.ParseUint(c.Param("id"), 10, 64)
	if iderr != nil {
		c.JSON(400, gin.H{
			"message": "id error",
		})
		return
	}

	var book entity.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	book.ID = bookId
	book, err = service.UpdateBookByID(book)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update book by id",
		"book":    book,
	})
}

func DeleteBookByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := service.DeleteBookByID(id)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete book by id",
	})
}
