package books

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"time"
)

// Book contains all Book's table properties.
type Book struct {
	ID             uint       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title          string     `json:"title"`
	Author         string     `json:"author"`
	Description    string     `json:"description"`
	Image          string     `json:"image"`
	RegisterNumber string     `json:"registerNumber" gorm:"unique"`
	Available      bool       `json:"available" gorm:"default:true"`
	CategoriesId   string     `json:"categoriesId,omitempty" gorm:"->"`                                       // Read Only
	BookCategory   []Category `gorm:"one2many:bookCategory,->;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` // Read Only
	CreatedAt      time.Time  `time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time  `time_format:"2006-01-02 15:04:05"`
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

type Category struct {
	ID         uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID     uint `json:"bookId" gorm:"not null"`
	CategoryID uint `json:"categoryId" gorm:"not null"`
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
