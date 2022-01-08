package books

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"gorm.io/gorm"
)

// Book contains all Book's table properties.
type Book struct {
	gorm.Model
	Title          string           `json:"title"`
	Author         string           `json:"author"`
	Description    string           `json:"description"`
	Image          string           `json:"image"`
	RegisterNumber string           `json:"registerNumber" gorm:"unique"`
	Available      bool             `json:"available" gorm:"default:true"`
	CategoriesId   string           `json:"categoriesId,omitempty" gorm:"->"` // Read Only
	BookCategories []BookCategories `gorm:"->"`
}

func (b Book) TableName() string {
	return "books"
}

type Module interface {
	Get() ([]Book, *errors.ApiError)
	GetByFilter(filter Filter) ([]Book, *errors.ApiError)
	Find(id string) (Book, *errors.ApiError)
	Create(book Book) (*Book, *errors.ApiError)
	Update(id string, upBook Book) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() (books []Book, apiError *errors.ApiError)
	GetByFilter(filter Filter) (books []Book, apiError *errors.ApiError)
	Find(id string) (book Book, apiError *errors.ApiError)
	Create(book Book) (*Book, *errors.ApiError)
	Update(id string, upBook Book) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}

type BookCategories struct {
	gorm.Model
	BookID     uint `json:"bookId" gorm:"not null"`
	CategoryID uint `json:"categoryId" gorm:"not null"`
}

func (bc BookCategories) TableName() string {
	return "book_categories"
}

type CategoryRepository interface {
	CreateCategories(bookId uint, categoriesIds []uint)
	DeleteCategories(bookId uint)
}

type Filter struct {
	Title      string `json:"title,omitempty" column:"books.title" array:"false"`
	Author     string `json:"author,omitempty" column:"books.author" array:"false"`
	Available  bool   `form:"available" column:"books.available" array:"false"`
	Categories string `form:"categ,omitempty" column:"book_categories.category_id" array:"true"`
}

func (f Filter) GetFieldNames() []string {
	return []string{"Title", "Author", "Categories", "Available"}
}
