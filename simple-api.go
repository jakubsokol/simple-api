package main

import  (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Person struct {
	ID uint `json:"id'`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func main() {
	r := gin.Default()
	r.GET("/people", getPeople)
	r.GET("/people/:id", getPersonById)
	r.POST("/people", createPerson)
	r.DELETE("people/:id", deletePerson)
	r.Run(":8181")
}

func getPersonById(context *gin.Context) {
	id := context.Params.ByName("id")
	person, err := getPersonByIdFromDb(id)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(404)
	} else {
		context.JSON(200, &person)
	}
}

func getPeople(context *gin.Context) {
	people, err := getAllPeopleFromDb()
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatus(404)
	} else {
		context.JSON(200, &people)
	}

}

func createPerson(context *gin.Context) {
	var person Person
	context.BindJSON(&person)
	err := addPersonToDb(person)
	if err == nil {
		context.JSON(200, person)
	} else {
		fmt.Println(err)
		context.AbortWithStatus(404)
	}
}

func deletePerson(context *gin.Context) {
	id := context.Params.ByName("id")
	err := deletePersonFromDb(id)
	if err == nil {
		context.JSON(200, gin.H{"id #" + id: "deleted"})
	} else {
		fmt.Println(err)
		context.AbortWithStatus(404)
	}
}


