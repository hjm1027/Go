package main

import "github.com/gin-gonic/gin"

type RegisterPayload struct {
    Username  string  `json:"username"`
    Password  string  `json:"password"`
}

func main() {
    router := gin.Default()
    router.POST("/register", func(c *gin.Context) {
        var data RegisterPayload
        if err := c.BindJSON(&data); err != nil {
            c.JSON(400, gin.H{
                "message": "Bad Request",
            })
            return
        }
        c.JSON(200, gin.H{
            "result": data.Username + data.Password,
        })
        return
    })

	router.Run(":8100")
}
