package schema

import (
	"fmt"
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Env contains an instance of MYSQL DB
type Env struct {
	db  *gorm.DB
	dsn string
}

// create database.
func createDatabase() (env *Env) {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASS"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS Library")
	db.Exec("USE Library")

	return &Env{db: db, dsn: dsn}
}

func (env *Env) createTables() {
	env.db.Migrator().AutoMigrate(&models.Student{})
	env.db.Migrator().AutoMigrate(&models.Book{})
	env.db.Migrator().AutoMigrate(&models.Category{})
	env.db.Migrator().AutoMigrate(&models.Lending{})
}

// Create create database and tables if notmysql exists
func Create() *gorm.DB {
	env := createDatabase()
	env.createTables()

	models.DB = env.db

	return env.db
}
