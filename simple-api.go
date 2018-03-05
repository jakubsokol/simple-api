package main

import  (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Person struct {
	ID uint `json:"id'`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var db *gorm.DB

func main() {
	db, err := gorm.Open("sqlite3", "/home/kuba/go/src/simple-api/gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/people", getPeople)
	r.GET("/people/:id", getPersonById)
	r.POST("/people", createPerson)

	r.Run(":8181")
}

func getPersonById(context *gin.Context) {
	id := context.Params.ByName("id")
	var person Person
	db, err := gorm.Open("sqlite3", "/home/kuba/go/src/simple-api/gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		context.JSON(200, person)
	}
}

func getPeople(context *gin.Context) {
	var people []Person
	db, err := gorm.Open("sqlite3", "/home/kuba/go/src/simple-api/gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	if err := db.Find(&people).Error; err != nil {
		context.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		context.JSON(200, people)
	}
}

func createPerson(context *gin.Context) {
	var person Person
	context.BindJSON(&person)
	db, err := gorm.Open("sqlite3", "/home/kuba/go/src/simple-api/gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&person)
	context.JSON(200, person)
}


