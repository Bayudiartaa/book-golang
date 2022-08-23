package books

import (
	"github.com/Bayudiartaa/book-golang/models"
	"gorm.io/gorm"
)

func FindBook(db *gorm.DB, book *models.Book, id int) error {
	return db.First(book, id).Error
}
