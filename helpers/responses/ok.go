package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Bayudiartaa/book-golang/models"
)

func Ok(c *gin.Context, books *[]models.Book, totalData int) {
	c.JSON(http.StatusOK, gin.H{
		"data":      books,
		"totalData": totalData,
	})
}
