package books

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Bayudiartaa/book-golang/helpers/responses"
	"github.com/Bayudiartaa/book-golang/models"
	booksServices "github.com/Bayudiartaa/book-golang/services/books"
)

// GET /books
// Find all books
func (repo *BooksRepository) FindAll(c *gin.Context) {
	var books []models.Book

	page := c.Query("page")
	limit := c.Query("limit")

	// search by query
	titleQuery := c.Query("title")
	authorQuery := c.Query("author")
	isbnQuery := c.Query("isbn")
	publisherQuery := c.Query("publisher")

	if titleQuery != "" || authorQuery != "" || isbnQuery != "" || publisherQuery != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			responses.BadRequest(c, err, "Failed to convert page to int")
			return
		}

		limit, err := strconv.Atoi(limit)
		if err != nil {
			responses.BadRequest(c, err, "Failed to convert page to int")
			return
		}

		if titleQuery != "" {
			totalData, err := booksServices.GetTotalDataByQuery(repo.DB, "title", titleQuery)
			if err != nil {
				responses.InternalServerError(c, err)
				return
			}
			if err := booksServices.FindAllBooksByQueryWithPagination(repo.DB, &books, "title", titleQuery, page, limit); err != nil {
				responses.BadRequest(c, err, "Failed to query books by title")
				return
			}
			responses.Ok(c, &books, totalData)
			return
		}

		if authorQuery != "" {
			totalData, err := booksServices.GetTotalDataByQuery(repo.DB, "author", authorQuery)
			if err != nil {
				responses.InternalServerError(c, err)
				return
			}
			if err := booksServices.FindAllBooksByQueryWithPagination(repo.DB, &books, "author", authorQuery, page, limit); err != nil {
				responses.BadRequest(c, err, "Failed to query books by author")
				return
			}
			responses.Ok(c, &books, totalData)
			return
		}

		if isbnQuery != "" {
			totalData, err := booksServices.GetTotalDataByQuery(repo.DB, "isbn", isbnQuery)
			if err != nil {
				responses.InternalServerError(c, err)
				return
			}
			if err := booksServices.FindAllBooksByQueryWithPagination(repo.DB, &books, "isbn", isbnQuery, page, limit); err != nil {
				responses.BadRequest(c, err, "Failed to query books by isbn")
				return
			}
			responses.Ok(c, &books, totalData)
			return
		}

		if publisherQuery != "" {
			totalData, err := booksServices.GetTotalDataByQuery(repo.DB, "publisher", publisherQuery)
			if err != nil {
				responses.InternalServerError(c, err)
				return
			}
			if err := booksServices.FindAllBooksByQueryWithPagination(repo.DB, &books, "publisher", publisherQuery, page, limit); err != nil {
				responses.BadRequest(c, err, "Failed to query books by publisher")
				return
			}
			responses.Ok(c, &books, totalData)
			return
		}
	}
	// ==============

	if page != "" && limit != "" {
		page, err := strconv.Atoi(page)
		if err != nil {
			responses.BadRequest(c, err, "Failed to convert page to int")
			return
		}

		limit, err := strconv.Atoi(limit)
		if err != nil {
			responses.BadRequest(c, err, "Failed to convert page to int")
			return
		}

		if err := booksServices.FindAllBooksWithLimit(repo.DB, &books, page, limit); err != nil {
			responses.InternalServerError(c, err)
			return
		}

		fmt.Println("Paginated books printed!")
		totalData, err := booksServices.GetTotalData(repo.DB)
		if err != nil {
			responses.InternalServerError(c, err)
			return
		}

		responses.Ok(c, &books, totalData)
		return
	}

	if err := booksServices.FindAllBooksWithoutLimit(repo.DB, &books); err != nil {
		responses.InternalServerError(c, err)
		return
	}

	fmt.Println("All books printed!")
	responses.Ok(c, &books, len(books))
}
