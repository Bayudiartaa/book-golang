package books

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Bayudiartaa/book-golang/models"
	booksServices "github.com/Bayudiartaa/book-golang/services/books"
)

type UpdateBookInput struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	Description   string `json:"description"`
	Publisher     string `json:"publisher"`
	NumberOfPages uint   `json:"numberOfPages"`
	CoverImage    string `json:"coverImage"`
}

// PATCH /books/:id
// Update book
func (repo BooksRepository) UpdateBook(c *gin.Context) {
	// get model if exist
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

	// validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to bind json",
		})
		return
	}

	var coverImage string
	if input.CoverImage == "" {
		coverImage = book.CoverImage
	} else {
		coverImage = input.CoverImage
	}

	updatedBook := models.Book{
		Title:         input.Title,
		Author:        input.Author,
		ISBN:          input.ISBN,
		Description:   input.Description,
		Publisher:     input.Publisher,
		NumberOfPages: input.NumberOfPages,
		CoverImage:    coverImage,
	}

	if err := booksServices.UpdateBook(repo.DB, &book, updatedBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
