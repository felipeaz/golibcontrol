package books

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
	"time"
)

// Book contains all Book's table properties.
type Book struct {
	ID              uint                  `json:"id" gorm:"primaryKey"`
	Title           string                `json:"title"`
	Author          string                `json:"author"`
	Description     string                `json:"description"`
	Image           string                `json:"image"`
	Available       bool                  `json:"available" gorm:"default:false"`
	CategoriesId    string                `json:"categoriesId,omitempty" gorm:"->"` // Read Only
	Registry        []registries.Registry `json:"registries" gorm:"->"`
	BookCategories  []BookCategories      `json:"categories" gorm:"->"`
	AvailableCopies int                   `json:"availableCopies" gorm:"->"`
	CreatedAt       time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
}

func (b Book) TableName() string {
	return "books"
}

func (b *Book) AfterFind(tx *gorm.DB) error {
	copies := 0
	var isAvailable bool
	for _, c := range b.Registry {
		if c.Available {
			copies++
			isAvailable = true
		}
	}
	b.AvailableCopies = copies
	b.Available = isAvailable
	return nil
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
	Get() ([]Book, *errors.ApiError)
	GetByFilter(filter Filter) ([]Book, *errors.ApiError)
	Find(id string) (Book, *errors.ApiError)
	Create(book Book) (*Book, *errors.ApiError)
	Update(id string, upBook Book) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type BookCategories struct {
	ID         uint                `json:"id" gorm:"primaryKey"`
	BookID     uint                `json:"bookId" gorm:"not null"`
	CategoryID uint                `json:"categoryId" gorm:"not null"`
	Category   categories.Category `json:"info" gorm:"->"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
}

func (bc BookCategories) TableName() string {
	return "book_categories"
}

type CategoryRepository interface {
	CreateCategories(bookId uint, categoriesIds []uint)
	DeleteCategories(bookId uint)
}

type Filter struct {
	Title          string `json:"title,omitempty" column:"books.title" array:"false" like:"true"`
	Author         string `json:"author,omitempty" column:"books.author" array:"false"`
	Available      bool   `form:"available" column:"books.available" array:"false"`
	Categories     string `form:"categ,omitempty" column:"book_categories.category_id" array:"true"`
	RegistryNumber string `form:"reg,omitempty" column:"book_registry_numbers.registry_number" array:"true"`
}

func (f Filter) GetFieldNames() []string {
	return []string{"Title", "Author", "Categories", "Available", "RegistryNumber"}
}
