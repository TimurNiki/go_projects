package main

import "github.com/gin-gonic/gin"

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Class   int    `json:"class"`
	Teacher string `json:"teacher"`
}

func listOfStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func main() {
	router := gin.Default()
	router.GET("/students", listOfStudents)
	router.Run("localhost:9595")

}
