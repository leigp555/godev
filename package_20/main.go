package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"package_pro/server"
)

func main()  {
	fmt.Println("hello")
	server.TestServer()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"x": " <div style={color:red}>lgp home</div>",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
