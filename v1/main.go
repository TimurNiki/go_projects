// package v1

// import (
// 	"errors"
// 	"strconv"

	
// )

// var students = []Student{
// 	{ID: 1, Name: "John", Class: 10, Teacher: "Mr. X"},
// 	{ID: 2, Name: "Jane", Class: 11, Teacher: "Mr. Y"},
// 	{ID: 3, Name: "Joe", Class: 12, Teacher: "Mr. Z"},
// }

// func listOfStudents(c *gin.Context) {
// 	c.JSON(200, students)
// }

// func createStudent(c *gin.Context) {
// 	var studentByUser Student
// 	//* Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// 	err := c.BindJSON(&studentByUser)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid request body"})
// 		return
// 	}
// 	students = append(students, studentByUser)
// 	c.JSON(200, gin.H{"message": "Student created successfully"})
// }

// //* first test
// // func getStudentByID(c *gin.Context) {
// // 	str_id := c.Param("id")
// // 	id, err := strconv.Atoi(str_id)
// // 	if err != nil {
// // 		c.JSON(400, gin.H{"error": "Invalid student ID"})
// // 		return
// // 	}
// // 	for _, student := range students {
// // 		if student.ID == id {
// // 			c.JSON(200, student)
// // 			return
// // 		}
// // 	}
// // 	c.JSON(404, gin.H{"error": "Student not found"})
// // }

// func main() {
// 	router := gin.Default()
// 	router.GET("/students", listOfStudents)
// 	router.GET("/students/:id", getStudent)
// 	router.POST("/students", createStudent)
// 	router.Run("localhost:9595")

// }

// func getStudentBI(i_id int) (*Student, error) {
// 	for i, s := range students {
// 		if s.ID == i_id {
// 			return &students[i], nil
// 		}
// 	}
// 	return nil, errors.New("Student not found")
// }
// func getStudent(c *gin.Context) {
// 	s_id := c.Param("id")
// 	i_id, err := strconv.Atoi(s_id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	student, err := getStudentBI(i_id)
// 	if err != nil {
// 		c.JSON(404, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, student)
// }


package main

import "fmt"

func main(){
	fmt.Println(`0`)
}