package database

import (
	"fmt"
	commentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	conferenceModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/conferences"
	groupModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	replyModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	reserveModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	reviewModel "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBHandler contains the connection with mysql that will be used on repositories layer
type DBHandler struct {
	conn *gorm.DB
}

// Connect opens a connection to the MYSQL server and prapare tables.
func Connect(user, password, host, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

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

// autoMigrateTables creates tables based on constants defined on internal if those tables doesn't exist.
// If the tables exists, this function will check if all properties of the structs are set on the tables
// and if the properties aren't set, updates them just like the struct definition.
func (db *DBHandler) autoMigrateTables() error {
	return db.conn.Migrator().AutoMigrate(
		&commentModel.Comment{},
		&replyModel.Reply{},
		&reserveModel.Reserve{},
		&reviewModel.Review{},
		&conferenceModel.Conference{},
		&conferenceModel.ConferenceSubscribers{},
		&groupModel.Group{},
		&groupModel.GroupSubscribers{},
	)
}
