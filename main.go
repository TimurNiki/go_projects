package main

import "github.com/gin-gonic/gin"

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   int    `json:"class"`
	Teacher string `json:"teacher"`
}

var students= []Student{
{ID:1,Name: "John",Class: 10,Teacher: "Mr. X"},
{ID:2,Name: "Jane",Class: 11,Teacher: "Mr. Y"},
{ID:3,Name: "Joe",Class: 12,Teacher: "Mr. Z"},
}

func listOfStudents(c *gin.Context) {
	c.JSON(200, students)
}

func createStudent(c *gin.Context){

}

func main() {
	router := gin.Default()
	router.GET("/students", listOfStudents)
	router.POST("/students", createStudent)
	router.Run("localhost:9595")

}
