package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message": "HelloWorld",
	})
}

func main() {
	fmt.Println("HelloWorld")
	r := gin.Default()
	r.GET("/", HomePage)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}