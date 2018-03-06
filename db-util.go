package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

const DB_FILE = "/home/kuba/go/src/simple-api/gorm.db"
const DB_ENGINE = "sqlite3"
var db *gorm.DB

func getAllPeopleFromDb() ([]Person, error) {
	var people []Person
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		err := db.Find(&people).Error
		return people, err
	} else {
		fmt.Println("Error while connect to DB", err)
		return nil, err
	}
}

func getPersonByIdFromDb(id string) (Person, error) {
	var person Person
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		err := db.Where("id = ?", id).First(&person).Error
		return person, err
	} else {
		fmt.Println("Error while connect to DB", err)
		return person, err
	}
}

func addPersonToDb(person Person) error{
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		db.Create(&person)
		return err
	} else {
		fmt.Println("Error while connect to DB", err)
		return err
	}
}
