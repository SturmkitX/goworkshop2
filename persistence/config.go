package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"goworkshop/model"
)

const (
	UNIQUE_BOOK_TITLE_CONSTRAINT = "title_unique"
)

var DBInstance *gorm.DB

func GetDB() (*gorm.DB, error) {
	if DBInstance == nil {
		DBInstance, err := InitDB()
		return DBInstance, err
	}

	return DBInstance, nil
}

func InitDB() (*gorm.DB, error) {

	DBInstance, err := gorm.Open("postgres", "host=localhost port=5432 user=db_admin " +
		"password=db_admin dbname=workshop_db sslmode=disable")

	if err != nil {
		fmt.Printf("Error while aquiring db connection: %s", err)
		return nil, err
	}
	DBInstance.DB().SetMaxOpenConns(20)

	DBInstance.LogMode(true)

	// Tables should be singular
	DBInstance.SingularTable(true)

	// Migrating the schema
	// This call will only create the new table if it does not exist - or add new columns , it will not modify
	// any data present in the table or remove or modify any columns
	DBInstance.AutoMigrate(model.Author{})
	DBInstance.AutoMigrate(model.Book{})

	// Adding foreign key constraints on the book table
	DBInstance.Table("book").AddForeignKey("author_uuid", "author(uuid)", "RESTRICT", "RESTRICT")

	//add uniqueness on book.title column
	DBInstance.Table("book").AddUniqueIndex(UNIQUE_BOOK_TITLE_CONSTRAINT, "title")

	return DBInstance, nil
}
