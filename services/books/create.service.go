package books

import (
	"github.com/Bayudiartaa/book-golang/models"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB, book *models.Book) error {
	return db.Create(book).Error
}
