package books

import (
	"github.com/Bayudiartaa/book-golang/models"
	"gorm.io/gorm"
)

func DeleteBook(db *gorm.DB, book *models.Book) error {
	return db.Delete(book).Error
}
