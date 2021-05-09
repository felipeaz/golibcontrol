package mysql

import (
	"fmt"
	"log"
	"os"

	accountModel "github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
	categoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/lending/model"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/student/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBHandler contains the connection with mysql that will be used on repositories layer
type DBHandler struct {
	conn *gorm.DB
}

// Connect opens a connection to the MYSQL server and prapare tables.
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbHandler := &DBHandler{conn: db}
	err = dbHandler.autoMigrateTables()

	return dbHandler.conn, err
}

// CloseConnection closes the connection to the MYSQL server.
func CloseConnection(db *gorm.DB) {
	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sql.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// autoMigrateTables creates tables based on constants defined on internal if those tables doesn't exists.
// If the tables exists, this function will check if all properties of the structs are set on the tables
// and if the properties aren't set, updates them just like the struct definition.
func (db *DBHandler) autoMigrateTables() error {
	return db.conn.Migrator().AutoMigrate(
		&studentModel.Student{},
		&bookModel.Book{},
		&categoryModel.Category{},
		&bookModel.BookCategory{},
		&lendingModel.Lending{},
		&accountModel.Account{},
	)
}
