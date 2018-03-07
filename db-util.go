package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

const DB_FILE = "/home/kuba/go/src/simple-api/gorm.db"
const DB_ENGINE = "sqlite3"

func getAllPeopleFromDb() ([]Person, error) {
	var people []Person
	var err error
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		err = db.Find(&people).Error
	} else {
		fmt.Println("Error while connect to DB", err)
	}
	return people, err
}

func getPersonByIdFromDb(id string) (Person, error) {
	var person Person
	var err error
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		err = db.Where("id = ?", id).First(&person).Error
	} else {
		fmt.Println("Error while connect to DB", err)
	}
	return person, err
}

func addPersonToDb(person Person) error {
	var err error
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		db.Create(&person)
	} else {
		fmt.Println("Error while connect to DB", err)
	}
	return err
}

func deletePersonFromDb(id string) error {
	var person Person
	var err error
	db, err := gorm.Open(DB_ENGINE, DB_FILE)
	db.AutoMigrate(&Person{})
	defer db.Close()
	if err == nil {
		err = db.Where("id = ?", id).Delete(&person).Error
	} else {
		fmt.Println("Error while connect to DB", err)
	}
	return err
}
