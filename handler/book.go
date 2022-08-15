package handler

import (
	"fmt"
	"net/http"

	"go-crud/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Dzaki Abdurrahman",
		"bio":  "Software Engineer in CODE ID",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "Hello Dzaki Abdurrahman",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title": title})
}

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	// if err != nil {
	// 	// log.Fatal(err)
	// 	for _, e := range err.(validator.ValidationErrors) {
	// 		errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
	// 		c.JSON(http.StatusBadRequest, errorMessage)
	// 		return
	// 	}
	// }
	if err != nil {
		//error message untuk 2 error
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
