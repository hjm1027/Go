package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, gin!",
        })
    })

	router.Run(":8100")
}
