package books

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Bayudiartaa/book-golang/models"
	booksServices "github.com/Bayudiartaa/book-golang/services/books"
)

// DELETE /books/:id
func (repo BooksRepository) DeleteBook(c *gin.Context) {
	var book models.Book
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to convert id to int",
		})
		return
	}

	if err := booksServices.FindBook(repo.DB, &book, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Book not found",
		})
		return
	}

	if err := booksServices.DeleteBook(repo.DB, &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v deleted successfully", c.Param("id")),
	})
}
