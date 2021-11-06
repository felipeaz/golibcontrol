package _interface

type BookCategoryRepositoryInterface interface {
	CreateCategories(bookId uint, categoriesIds []uint)
	DeleteCategories(bookId uint)
}
