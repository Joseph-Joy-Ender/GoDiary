package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func homePage(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func postHomePage(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Post Home page",
	})
}
func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", homePage)
	r.POST("/", postHomePage)
	r.GET("/query", queryString)

	err := r.Run()
	if err != nil {
		return
	}

}

func queryString(context *gin.Context) {
	name := context.Query("name")
	age := context.Query("age")

	context.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})

}
