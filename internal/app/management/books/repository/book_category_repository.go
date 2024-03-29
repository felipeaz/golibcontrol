package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"log"
	"strconv"
)

type BookCategoryRepository struct {
	DB database.GORMServiceInterface
}

func NewBookCategoryRepository(db database.GORMServiceInterface) BookCategoryRepository {
	return BookCategoryRepository{
		DB: db,
	}
}

// CreateCategories persists category on DB if exists.
func (r BookCategoryRepository) CreateCategories(bookId uint, categoriesIds []uint) {
	if len(categoriesIds) <= 0 {
		return
	}
	for i := 0; i < len(categoriesIds); i++ {
		bookCategory := books.BookCategories{
			BookID:     bookId,
			CategoryID: categoriesIds[i],
		}
		err := r.DB.Persist(&bookCategory)
		if err != nil {
			log.Println(errors.FailedToCreateBookCategoryMessage, err.Error())
		}
	}
}

// DeleteCategories removes a Book categories from DB
func (r BookCategoryRepository) DeleteCategories(bookId uint) {
	tx := r.DB.Where(nil, fmt.Sprintf("book_id = %s", strconv.Itoa(int(bookId))))
	err := r.DB.Remove(tx, books.BookCategories{})
	if err != nil {
		log.Println(err.Error())
	}
}
