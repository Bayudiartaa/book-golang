package books

import (
	"github.com/Bayudiartaa/book-golang/models"
	"gorm.io/gorm"
)

func UpdateBook(db *gorm.DB, book *models.Book, updatedBook models.Book) error {
	return db.Model(book).Updates(updatedBook).Error
}
